package importer

import (
	"context"
	"fmt"
	"time"

	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/cassandra"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/wikipedia"
	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"go.uber.org/zap"
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

// TODO: check if the leaving out of event attributes except day and month will not send it in the repsonse
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
	return fmt.Sprintf("failed to import %s event at %s stage on %d-%d: %s", e.Type, e.Stage ,e.Month, e.Day, e.Err)
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

// Import implements Interface.
func (i *Impl) Import(ctx context.Context) ImportResult {
	panic("unimplemented")
}
