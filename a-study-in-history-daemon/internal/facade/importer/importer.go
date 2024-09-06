package importer

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/cassandra"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/wikipedia"
	"github.com/aaegamysta/a-study-in-history/daemon/pkg/tiime"
	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type (
	ImportStatus int64
	Stage        int64
)

const (
	Success ImportStatus = iota
	PartialSuccess
	Failed
)

const (
	Fetching Stage = iota
	Persisting
)

func (s Stage) String() string {
	switch s {
	case Fetching:
		return "Fetching"
	case Persisting:
		return "Persisting"
	default:
		panic("unknown stage")
	}
}

// TODO: check if the leaving out of event attributes except day and month will not send it in the response.
type ImportResult struct {
	Status          ImportStatus
	MissedOutEvents []events.Event
	ImportedOn      time.Time
}

type ImportPipelineStageResult struct {
	Events events.EventsCollection
	Error  error
}

type PipelineStageError struct {
	Stage Stage
	Type  events.Type
	Month int64
	Day   int64
	Err   error
}

func (e PipelineStageError) Error() string {
	return fmt.Sprintf("failed to import %s event at %s stage on %d-%d: %s", e.Type, e.Stage, e.Month, e.Day, e.Err)
}

type Interface interface {
	Import(ctx context.Context) ImportResult
}

type Impl struct {
	cfg             Config
	logger          *zap.SugaredLogger
	cassandraClient cassandra.Interface
	wikipediaClient wikipedia.Interface
}

func New(ctx context.Context, logger *zap.SugaredLogger, cfg Config,
	cassandraClient cassandra.Interface, wikipediaClient wikipedia.Interface,
) Interface {
	impl := &Impl{
		logger:          logger,
		cfg:             cfg,
		cassandraClient: cassandraClient,
		wikipediaClient: wikipediaClient,
	}
	if cfg.ImportAtStart {
		importResult := impl.Import(ctx)
		impl.logger.Debugf("importing at start complete at %s was %s with %s missed events", importResult.ImportedOn, importResult.Status, len(importResult.MissedOutEvents))
	}
	return impl
}

func (i *Impl) Import(ctx context.Context) ImportResult {
	if i.cfg.ImportConcurrently {
		result := i.importConcurrently(ctx)
		return result
	}
	result := i.importSequentially(ctx)
	return result
}

func (i *Impl) importConcurrently(ctx context.Context) ImportResult {
	eg, ctx := errgroup.WithContext(ctx)
	var historicalEventsImportResult ImportResult
	var birthsImportResult ImportResult
	var deathsImportResult ImportResult
	var holidaysImportResult ImportResult
	eg.Go(func() error {
		historicalEventsImportResult = i.doImport(ctx, events.Historical)
		return nil
	})
	eg.Go(func() error {
		birthsImportResult = i.doImport(ctx, events.Birth)
		return nil
	})
	eg.Go(func() error {
		deathsImportResult = i.doImport(ctx, events.Death)
		return nil
	})
	eg.Go(func() error {
		holidaysImportResult = i.doImport(ctx, events.Holiday)
		return nil
	})
	err := eg.Wait()
	if err != nil {
	}
	aggregatedImportResult := aggregateImportResults(historicalEventsImportResult,
		birthsImportResult,
		deathsImportResult,
		holidaysImportResult,
	)
	return aggregatedImportResult
}

func (i *Impl) importSequentially(ctx context.Context) ImportResult {
	historicalEventsImportResult := i.doImport(ctx, events.Historical)
	birthsImportResult := i.doImport(ctx, events.Birth)
	deathsImportResult := i.doImport(ctx, events.Death)
	holidaysImportResult := i.doImport(ctx, events.Holiday)
	aggregatedImportResult := aggregateImportResults(historicalEventsImportResult, birthsImportResult, holidaysImportResult, deathsImportResult)
	return aggregatedImportResult
}

func (i *Impl) doImport(ctx context.Context, typing events.Type) ImportResult {
	eventsFinder := make([]<-chan ImportPipelineStageResult, 365)
	for dayOfYear := range tiime.DaysInYear {
		month, day := tiime.DayOfYearToMonthDay(int64(dayOfYear) + 1)
		eventsFinder[dayOfYear] = i.persistEventsPipelineStage(ctx, i.retrieveEventsPipelineStage(ctx, typing, month, day))
	}
	importResult := i.fanInFetchedPersistedEvents(ctx, eventsFinder...)
	return importResult
}

func (i *Impl) retrieveEventsPipelineStage(ctx context.Context, typing events.Type, month, day int64) <-chan ImportPipelineStageResult {
	retrievedEventsStream := make(chan ImportPipelineStageResult)
	go func() {
		defer close(retrievedEventsStream)
		select {
		case <-ctx.Done():
			return
		default:
			var coll wikipedia.EventsCollectionResult
			var err error
			switch typing {
			case events.Historical:
				coll, err = i.wikipediaClient.ImportHistoricalEventsFor(ctx, month, day)
			case events.Birth:
				coll, err = i.wikipediaClient.ImportBirthsFor(ctx, month, day)
			case events.Death:
				coll, err = i.wikipediaClient.ImportDeathsFor(ctx, month, day)
			case events.Holiday:
				coll, err = i.wikipediaClient.ImportHolidaysFor(ctx, month, day)
			}
			if err != nil {
				retrievedEventsStream <- ImportPipelineStageResult{
					Error: PipelineStageError{
						Stage: Fetching,
						Type:  typing,
						Month: month,
						Day:   day,
					},
				}
				return
			}
			retrievedEventsStream <- ImportPipelineStageResult{
				Events: coll.Coll,
				Error:  nil,
			}
		}
	}()
	return retrievedEventsStream
}

func (i *Impl) persistEventsPipelineStage(ctx context.Context, retrievedEventsStream <-chan ImportPipelineStageResult) <-chan ImportPipelineStageResult {
	persistedEventsStream := make(chan ImportPipelineStageResult)
	go func() {
		defer close(persistedEventsStream)
		select {
		case <-ctx.Done():
			return
		case eventsColl := <-retrievedEventsStream:
			if eventsColl.Error != nil {
				persistedEventsStream <- eventsColl
			}
			err := i.cassandraClient.UpsertEvents(ctx, eventsColl.Events)
			if err != nil {
				persistedEventsStream <- ImportPipelineStageResult{
					Events: events.EventsCollection{},
					Error: PipelineStageError{
						Stage: Persisting,
						Type:  eventsColl.Events.Type,
						Month: eventsColl.Events.Month,
						Day:   eventsColl.Events.Day,
						Err:   err,
					},
				}
				return
			}
			persistedEventsStream <- ImportPipelineStageResult{
				Events: eventsColl.Events,
				Error:  nil,
			}
		}
	}()
	return persistedEventsStream
}

func (i *Impl) fanInFetchedPersistedEvents(ctx context.Context, eventsStream ...<-chan ImportPipelineStageResult) ImportResult {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	// TODO: see if using atomic counter for totalFetchedEvents would be better instead of having in critical section
	missedOutEvents := make([]events.Event, 0)
	totalFetchedEvents := 0

	multiplex := func(ch <-chan ImportPipelineStageResult) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case result := <-ch:
			mu.Lock()
			defer mu.Unlock()
			if result.Error != nil {
				stageErr := result.Error.(PipelineStageError)
				missedOutEvents = append(missedOutEvents, events.Event{
					Type:  stageErr.Type,
					Day:   stageErr.Day,
					Month: stageErr.Month,
				})
				i.logger.Debugf("an error occurred at stage %s while importing %s events on %d-%d, missed out on %d events", stageErr.Type, stageErr.Type, stageErr.Month, stageErr.Day)
				return
			}
			totalFetchedEvents += len(result.Events.Collection)
			i.logger.Debugf("retrieved and persisted %s events on %d-%d totalling %d", result.Events.Type, result.Events.Month, result.Events.Day, len(result.Events.Collection))
		}
	}

	wg.Add(len(eventsStream))
	for _, stream := range eventsStream {
		go multiplex(stream)
	}

	wg.Wait()

	var importStatus ImportStatus
	switch {
	case len(missedOutEvents) == 0:
		importStatus = Success
	case len(missedOutEvents) < totalFetchedEvents:
		importStatus = PartialSuccess
	case totalFetchedEvents == 0:
		importStatus = Failed
	}
	i.logger.Debugf("fetching, persistence and fanning in of %d events with %d missed out events %s completed", totalFetchedEvents, len(missedOutEvents), importStatus)

	return ImportResult{
		Status:          importStatus,
		MissedOutEvents: missedOutEvents,
		ImportedOn:      time.Now(),
	}
}

func aggregateImportResults(historical, birth, death, holiday ImportResult) ImportResult {
	var aggregatedImportStatus ImportStatus
	switch {
	case historical.Status == Failed || birth.Status == Failed || death.Status == Failed || holiday.Status == Failed:
		aggregatedImportStatus = Failed
	case historical.Status == PartialSuccess || birth.Status == PartialSuccess || death.Status == PartialSuccess || holiday.Status == PartialSuccess:
		aggregatedImportStatus = PartialSuccess
	default:
		aggregatedImportStatus = Success
	}
	aggregatedMissedOutEvents := make([]events.Event, 0)
	slices.Concat(aggregatedMissedOutEvents, historical.MissedOutEvents,
		birth.MissedOutEvents,
		death.MissedOutEvents,
		holiday.MissedOutEvents,
	)
	return ImportResult{
		Status:          aggregatedImportStatus,
		MissedOutEvents: aggregatedMissedOutEvents,
		ImportedOn:      time.Now(),
	}
}
