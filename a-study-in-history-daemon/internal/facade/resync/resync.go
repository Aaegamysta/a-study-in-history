package resync

import (
	"context"
	//nolint: gosec // MD5 is used for generating unique IDs, not for cryptographic purposes
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/cassandra"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/wikipedia"
	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type TargetEvents int64

const (
	Unspecified TargetEvents = iota
	History
	Birth
	Death
	Holiday
	All
)

var ErrEventsOutOfSync = errors.New("events out of sync")

type Interface interface {
	Resynchronize(ctx context.Context, target TargetEvents, month, day int64) error
}

type Impl struct {
	logger          *zap.SugaredLogger
	wikipediaClient wikipedia.Interface
	cassandraClient cassandra.Interface
}

/*
 * TODO:
 * There might be an issue with this way of sorting for holidays as they do not have a specified year.
 */
func (i *Impl) Resynchronize(ctx context.Context, typo TargetEvents, month int64, day int64) error {
	eg, ctx := errgroup.WithContext(ctx)
	var totalFetchedEventsHash string
	var totalPersistedEventsHash string
	var candidateUpdatedEventsColl events.Collection
	var fromDataBase events.Collection
	eg.Go(func() error {
		var collRes wikipedia.EventsCollectionResult
		var err error
		switch typo {
		case History:
			collRes, err = i.wikipediaClient.ImportHistoricalEventsFor(ctx, month, day)
		case Birth:
			collRes, err = i.wikipediaClient.ImportBirthsFor(ctx, month, day)
		case Death:
			collRes, err = i.wikipediaClient.ImportDeathsFor(ctx, month, day)
		case Holiday:
			collRes, err = i.wikipediaClient.ImportHolidaysFor(ctx, month, day)
		case All:
			panic("add a method to the wikipedia client to fetch all events")
		default:
			return fmt.Errorf("unknown target events type: %d", typo)
		}
		if err != nil {
			return fmt.Errorf("something wrong happened while fetching events from wikipedia %w", err)
		}
		sort.Slice(collRes.Coll.Events, func(i, j int) bool {
			if collRes.Coll.Events[j].Year == collRes.Coll.Events[i].Year {
				return collRes.Coll.Events[j].ID < collRes.Coll.Events[i].ID
			}
			return collRes.Coll.Events[j].Year < collRes.Coll.Events[i].Year
		})
		//nolint: gosec // we are hashing the event ids for equality check, no sensitive data is being hashed
		h := md5.New()
		for _, e := range collRes.Coll.Events {
			h.Write([]byte(e.ID))
		}
		sum := h.Sum(nil)
		totalFetchedEventsHash = hex.EncodeToString(sum)
		candidateUpdatedEventsColl = collRes.Coll
		return nil
	})

	eg.Go(func() error {
		var coll events.Collection
		var err error
		switch typo {
		case History:
			coll, err = i.cassandraClient.ListHistoricalEventsFor(ctx, month, day)
		case Birth:
			coll, err = i.cassandraClient.ListBirthsFor(ctx, month, day)
		case Death:
			coll, err = i.cassandraClient.ListDeathsFor(ctx, month, day)
		case Holiday:
			coll, err = i.cassandraClient.ListHolidaysFor(ctx, month, day)
		case All:
			coll, err = i.cassandraClient.ListEventsFor(ctx, month, day)
		default:
			return fmt.Errorf("unknown target events type: %d", typo)
		}
		if err != nil {
			return fmt.Errorf("something wrong happened while fetching events from wikipedia: %d", typo)
		}
		sort.Slice(coll.Events, func(i, j int) bool {
			if coll.Events[j].Year == coll.Events[i].Year {
				return coll.Events[j].ID < coll.Events[i].ID
			}
			return coll.Events[j].Year < coll.Events[i].Year
		})
		//nolint: gosec // we are hashing the event ids for equality check, no sensitive data is being hashed
		h := md5.New()
		for _, e := range coll.Events {
			h.Write([]byte(e.ID))
		}
		sum := h.Sum(nil)
		fromDataBase = coll
		totalPersistedEventsHash = hex.EncodeToString(sum)
		return nil
	})

	err := eg.Wait()
	if err != nil {
		return fmt.Errorf("an error occurred while trying to upsert and update events during resynchronization: %w", err)
	}

	if totalFetchedEventsHash == totalPersistedEventsHash {
		i.logger.Info("events are in sync")
		return nil
	}
	i.logger.Infof("events are out of sync, updating events")
	err = i.cassandraClient.UpsertEvents(ctx, candidateUpdatedEventsColl)
	if err != nil {
		return fmt.Errorf("an error occurred while trying to upsert and update events during resynchronization: %w", err)
	}
	log.Println(fromDataBase.Events)
	return ErrEventsOutOfSync
}

func New(_ context.Context,
	logger *zap.SugaredLogger,
	wikipediaClient wikipedia.Interface,
	cassandraClient cassandra.Interface,
) Interface {
	return &Impl{
		logger:          logger,
		wikipediaClient: wikipediaClient,
		cassandraClient: cassandraClient,
	}
}
