package cassandra

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
	"sync"

	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Interface interface {
	CreateTablesIfNotExists(ctx context.Context) error
	ListHistoricalEventsFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	ListBirthsFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	ListDeathsFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	ListHolidaysFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	UpsertEvents(ctxc context.Context, coll events.EventsCollection) error
}

type connectionResult struct {
	conn *grpc.ClientConn
	err  error
}

type Impl struct {
	cfg            Config
	logger         *zap.SugaredLogger
	connectionPool sync.Pool
}

func New(ctx context.Context, cfg Config, logger *zap.SugaredLogger) Interface {
	impl := &Impl{
		cfg:    cfg,
		logger: logger,
	}
	impl.connectionPool = sync.Pool{
		New: func() any {
			config := &tls.Config{
				InsecureSkipVerify: false,
			}
			conn, err := grpc.NewClient(cfg.ConnectionString, grpc.WithTransportCredentials(credentials.NewTLS(config)),
				grpc.WithPerRPCCredentials(
					auth.NewStaticTokenProvider(cfg.Token),
				),
			)
			if err != nil {
				return &connectionResult{
					conn: nil,
					err: err,
				}
			}
			return &connectionResult{
				conn: conn,
				err:  nil,
			}
		},
	}
	return impl
}

// CreateTablesIfNotExists implements Interface.
func (i *Impl) CreateTablesIfNotExists(ctx context.Context) error {
	panic("unimplemented")
}

// ListBirthsFor implements Interface.
func (i *Impl) ListBirthsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	panic("unimplemented")
}

// ListDeathsFor implements Interface.
func (i *Impl) ListDeathsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	panic("unimplemented")
}

// ListHistoricalEventsFor implements Interface.
func (i *Impl) ListHistoricalEventsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	panic("unimplemented")
}

// ListHolidaysFor implements Interface.
func (i *Impl) ListHolidaysFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	panic("unimplemented")
}

// UpsertEvents implements Interface.
func (i *Impl) UpsertEvents(ctxc context.Context, coll events.EventsCollection) error {
	panic("unimplemented")
}
