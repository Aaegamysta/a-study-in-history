package server

import (
	"context"
	"errors"

	"github.com/aaegamysta/a-study-in-history/daemon/internal/facade/collector"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/facade/importer"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/facade/synch"
	"github.com/aaegamysta/a-study-in-history/spec/gen"
	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	gen.UnimplementedAStudyInHistoryServer
	logger   *zap.SugaredLogger
	importer importer.Interface
	synch    synch.Interface
	collect  collector.Interface
}

func New(_ context.Context, logger *zap.SugaredLogger, importer importer.Interface, synch synch.Interface, collect collector.Interface) gen.AStudyInHistoryServer {
	return Server{
		logger:   logger,
		importer: importer,
		synch:    synch,
		collect:  collect,
	}
}

// GetEventsFor implements gen.AStudyInHistoryServer.
func (s Server) GetEventsFor(ctx context.Context, req *gen.GetEventsRequest) (*gen.GetEventsResponse, error) {
	validationErr := validateDaysAndMonth(req.Month, req.Day)
	if validationErr != nil {
		return nil, validationErr
	}
	var coll events.EventsCollection
	var err error
	switch req.TargetEvents {
	case gen.TargetEvents_TARGET_EVENTS_BIRTH:
		coll, err = s.collect.ListBirthsFor(ctx, req.Month, req.Day)
	case gen.TargetEvents_TARGET_EVENTS_DEATH:
		coll, err = s.collect.ListDeathsFor(ctx, req.Month, req.Day)
	case gen.TargetEvents_TARGET_EVENTS_HISTORY:
		coll, err = s.collect.ListHistoricalEventsFor(ctx, req.Month, req.Day)
	case gen.TargetEvents_TARGET_EVENTS_HOLIDAY:
		coll, err = s.collect.ListHolidaysFor(ctx, req.Month, req.Day)
	case gen.TargetEvents_TARGET_EVENTS_UNSPECIFIED:
		fallthrough
	default:
		return nil, status.Errorf(codes.Internal, "an error happened while getting %s events for %d-%d", req.TargetEvents, req.Month, req.Day)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "an error happened while getting %s events for %d-%d", req.TargetEvents, req.Month, req.Day)
	}
	res := &gen.GetEventsResponse{
		Month:            coll.Month,
		Day:              coll.Day,
		EventsCollection: events.MapEventsCollectionToGRPC(coll),
	}
	return res, nil
}

// Import implements gen.AStudyInHistoryServer.
func (s Server) Import(ctx context.Context, _ *gen.ImportRequest) (*gen.ImportResponse, error) {
	result := s.importer.Import(ctx)
	missedEvents := make([]*gen.MissedEvents, 0)
	for _, e := range result.MissedOutEvents {
		missedEvents = append(missedEvents, &gen.MissedEvents{
			Type:  gen.Type(e.Type),
			Day:   e.Day,
			Month: e.Month,
		})
	}
	res := &gen.ImportResponse{
		Status:       gen.ImportStatus(result.Status),
		MissedEvents: missedEvents,
		ImportedOn:   result.ImportedOn.Unix(),
	}
	return res, nil
}

// SynchronizeFor implements gen.AStudyInHistoryServer.
func (s Server) SynchronizeFor(ctx context.Context, req *gen.SynchronizeForRequest) (*gen.SynchronizeForResponse, error) {
	err := validateDaysAndMonth(req.Month, req.Day)
	if err != nil {
		return nil, err
	}
	err = s.synch.Synchronize(ctx, synch.TargetEvents(req.TargetEvents), req.Month, req.Day)
	if errors.Is(err, synch.EventsOutOfSyncErr) {
		return &gen.SynchronizeForResponse{
			Status: true,
		}, nil
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something wrong happened while attempting to synchronize %s events for %d-%d", req.GetTargetEvents(), req.GetMonth(), req.GetDay())
	}
	return &gen.SynchronizeForResponse{
		Status: false,
	}, nil
}

func validateDaysAndMonth(month, day int64) error {
	if month < 1 || month > 12 {
		return status.Errorf(codes.InvalidArgument, "month must be between 1 and 12, got %d", month)
	}
	if day < 1 || day > 31 {
		return status.Errorf(codes.InvalidArgument, "day must be between 1 and 31, got %d", day)
	}
	return nil
}
