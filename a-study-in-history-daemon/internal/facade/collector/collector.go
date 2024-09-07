package collector

import (
	"context"

	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/cassandra"
	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"go.uber.org/zap"
)

type Interface interface {
	ListEventsFor(ctx context.Context, month, day int64) (events.Collection, error)
	ListHistoricalEventsFor(ctx context.Context, month, day int64) (events.Collection, error)
	ListBirthsFor(ctx context.Context, month, day int64) (events.Collection, error)
	ListDeathsFor(ctx context.Context, month, day int64) (events.Collection, error)
	ListHolidaysFor(ctx context.Context, month, day int64) (events.Collection, error)
}

type Implementation struct {
	logger     *zap.SugaredLogger
	repository cassandra.Interface
}

func New(_ context.Context, logger *zap.SugaredLogger, repository cassandra.Interface) Interface {
	return &Implementation{
		logger:     logger,
		repository: repository,
	}
}

// ListEventsFor implements Interface.
func (i *Implementation) ListEventsFor(ctx context.Context, month int64, day int64) (events.Collection, error) {
	coll, err := i.repository.ListEventsFor(ctx, month, day)
	if err != nil {
		i.logger.Errorw("failed to list all events", "error", err)
		return events.Collection{}, err
	}
	return coll, nil
}

// ListBirthsFor implements Interface.
func (i *Implementation) ListBirthsFor(ctx context.Context, month int64, day int64) (events.Collection, error) {
	coll, err := i.repository.ListBirthsFor(ctx, month, day)
	if err != nil {
		i.logger.Errorw("failed to list all events", "error", err)
		return events.Collection{}, err
	}
	return coll, nil
}

// ListDeathsFor implements Interface.
func (i *Implementation) ListDeathsFor(ctx context.Context, month int64, day int64) (events.Collection, error) {
	coll, err := i.repository.ListDeathsFor(ctx, month, day)
	if err != nil {
		i.logger.Errorw("failed to list all events", "error", err)
		return events.Collection{}, err
	}
	return coll, nil
}

// ListHistoricalEventsFor implements Interface.
func (i *Implementation) ListHistoricalEventsFor(ctx context.Context, month int64, day int64) (events.Collection, error) {
	coll, err := i.repository.ListHistoricalEventsFor(ctx, month, day)
	if err != nil {
		i.logger.Errorw("failed to list all events", "error", err)
		return events.Collection{}, err
	}
	return coll, nil
}

// ListHolidaysFor implements Interface.
func (i *Implementation) ListHolidaysFor(ctx context.Context, month int64, day int64) (events.Collection, error) {
	coll, err := i.repository.ListHolidaysFor(ctx, month, day)
	if err != nil {
		i.logger.Errorw("failed to list all events", "error", err)
		return events.Collection{}, err
	}
	return coll, nil
}
