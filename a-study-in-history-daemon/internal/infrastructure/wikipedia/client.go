package wikipedia

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"go.uber.org/zap"
)

type Interface interface {
	ImportHistoricalEventsFor(ctx context.Context, month, day int64) (EventsCollectionResult, error)
	ImportBirthsFor(ctx context.Context, month, day int64) (EventsCollectionResult, error)
	ImportDeathsFor(ctx context.Context, month, day int64) (EventsCollectionResult, error)
	ImportHolidaysFor(ctx context.Context, month, day int64) (EventsCollectionResult, error)
}

type Impl struct {
	logger     *zap.SugaredLogger
	httpClient http.Client
	path       string
}

type EventsCollectionResult struct {
	Coll events.EventsCollection
	Err  error
}

type EventsFetchError struct {
	Type   events.Type
	Day    int64
	Month  int64
	Reason error
}

func (e EventsFetchError) Error() string {
	return fmt.Sprintf("failed to fetch %s events on %d-%d: %v", e.Type, e.Month, e.Day, e.Reason)
}

func New(_ context.Context, cfg Config, logger *zap.SugaredLogger) Interface {
	return &Impl{
		logger:     logger,
		httpClient: http.Client{},
		path:       cfg.Path,
	}
}

// ImportHistoricalEventsFor implements Interface.
func (i *Impl) ImportHistoricalEventsFor(ctx context.Context, month int64, day int64) (EventsCollectionResult, error) {
	coll := events.EventsCollection{
		Type:   events.Historical,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	err := i.doImport(ctx, &coll)
	if err != nil {
		return EventsCollectionResult{}, &EventsFetchError{
			Type:   events.Historical,
			Day:    day,
			Month:  month,
			Reason: err,
		}
	}
	return EventsCollectionResult{
		Coll: coll,
		Err:  nil,
	}, nil
}

// ImportBirthsFor implements Interface.
func (i *Impl) ImportBirthsFor(ctx context.Context, month int64, day int64) (EventsCollectionResult, error) {
	coll := events.EventsCollection{
		Type:   events.Birth,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	err := i.doImport(ctx, &coll)
	if err != nil {
		return EventsCollectionResult{}, &EventsFetchError{
			Type:   events.Birth,
			Day:    day,
			Month:  month,
			Reason: err,
		}
	}
	return EventsCollectionResult{
		Coll: coll,
		Err:  nil,
	}, nil
}

// ImportDeathsFor implements Interface.
func (i *Impl) ImportDeathsFor(ctx context.Context, month int64, day int64) (EventsCollectionResult, error) {
	coll := events.EventsCollection{
		Type:   events.Death,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	err := i.doImport(ctx, &coll)
	if err != nil {
		return EventsCollectionResult{}, &EventsFetchError{
			Type:   events.Death,
			Day:    day,
			Month:  month,
			Reason: err,
		}
	}
	return EventsCollectionResult{
		Coll: coll,
		Err:  nil,
	}, nil
}

// ImportHolidaysFor implements Interface.
func (i *Impl) ImportHolidaysFor(ctx context.Context, month int64, day int64) (EventsCollectionResult, error) {
	coll := events.EventsCollection{
		Type:   events.Holiday,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	err := i.doImport(ctx, &coll)
	if err != nil {
		return EventsCollectionResult{}, &EventsFetchError{
			Type:   events.Holiday,
			Day:    day,
			Month:  month,
			Reason: err,
		}
	}
	return EventsCollectionResult{
		Coll: coll,
		Err:  nil,
	}, nil
}

func (i *Impl) doImport(ctx context.Context, coll *events.EventsCollection) error {
	url := fmt.Sprintf("%s/%s/%d/%d", i.path, coll.Type, coll.Month, coll.Day)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		i.logger.Panicf("failed to create request to fetch %s events for %d-%d: %v", coll.Type, coll.Month, coll.Day, err)
	}
	res, err := i.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("something wrong happened while importing %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("something wrong happened while importing, status code is not ok %w", err)
	}
	if res.Body == nil {
		return fmt.Errorf("something wrong happened while importing, response body is empty %w", err)
	}
	defer res.Body.Close()
	err = i.parseEvents(ctx, coll, res.Body)
	if err != nil {
		return fmt.Errorf("something wrong happened while parsing events %w", err)
	}
	return nil
}