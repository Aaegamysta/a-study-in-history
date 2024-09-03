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

func (i *Impl) ImportHistoricalEventsFor(ctx context.Context, month, day int64) (EventsCollectionResult, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Impl) ImportBirthsFor(ctx context.Context, month, day int64) (EventsCollectionResult, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Impl) ImportDeathsFor(ctx context.Context, month, day int64) (EventsCollectionResult, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Impl) ImportHolidaysFor(ctx context.Context, month, day int64) (EventsCollectionResult, error) {
	//TODO implement me
	panic("implement me")
}
