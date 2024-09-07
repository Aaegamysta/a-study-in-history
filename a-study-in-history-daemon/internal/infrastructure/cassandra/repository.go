package cassandra

import (
	"context"
	"crypto/tls"
	"fmt"
	"math"
	"slices"
	"sync"

	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Operation int64

const (
	Unspecified Operation = iota
	Upsert
	List
)

func (o Operation) String() string {
	switch o {
	case Unspecified:
		return "unspecified"
	case Upsert:
		return "upsert"
	case List:
		return "list"
	default:
		return "unknown"
	}
}

type Interface interface {
	CreateTablesIfNotExists(ctx context.Context) error
	ListEventsFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	ListHistoricalEventsFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	ListBirthsFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	ListDeathsFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	ListHolidaysFor(ctx context.Context, month, day int64) (events.EventsCollection, error)
	UpsertEvents(ctxc context.Context, coll events.EventsCollection) error
}

type Impl struct {
	cfg            Config
	logger         *zap.SugaredLogger
	connectionPool sync.Pool
}

type connectionResult struct {
	conn *grpc.ClientConn
	err  error
}

type DatabaseError struct {
	Operation Operation
	Type      events.Type
	Month     int64
	Day       int64
	Err       error
}

func (d DatabaseError) Error() string {
	return fmt.Sprintf("failed to %s %s events for %d-%d", d.Operation, d.Type, d.Month, d.Day)
}

func New(_ context.Context, cfg Config, logger *zap.SugaredLogger) Interface {
	impl := &Impl{
		cfg:    cfg,
		logger: logger,
	}
	impl.connectionPool = sync.Pool{
		New: func() any {
			//nolint:gosec // usgae of InsecureSkipVerify is intentional and as stated by documentation fo Datastax
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
					err:  err,
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
	connRes, ok := i.connectionPool.Get().(*connectionResult)
	if !ok {
		i.logger.Panicf("unexpected type while retrieving connection from pool")
	}
	if connRes.err != nil {
		return fmt.Errorf("something wrong happened while creating tables %w", connRes.err)
	}
	conn := connRes.conn
	// NewStargateClientWithConn implementation never returns an error
	cassandraClient, _ := client.NewStargateClientWithConn(conn)
	cql := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s.events_by_type_day_month (
			type text, day int, month int, year int, id text, title text ,description text, 
			thumbnail_source text, thumbnail_width int, thumbnail_height int,
			PRIMARY KEY ((type, day, month), year, id)
		)
	`, i.cfg.Keyspace)
	_, err := cassandraClient.ExecuteQueryWithContext(&proto.Query{
		Cql: cql,
	}, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i *Impl) ListEventsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	eg, ctx := errgroup.WithContext(ctx)

	var historicalEvents events.EventsCollection
	eg.Go(func() error {
		coll, err := i.ListBirthsFor(ctx, month, day)
		if err != nil {
			return err
		}
		historicalEvents = coll
		return nil
	})

	var birthEvents events.EventsCollection
	eg.Go(func() error {
		coll, err := i.ListBirthsFor(ctx, month, day)
		if err != nil {
			return err
		}
		birthEvents = coll
		return nil
	})

	var deathEvents events.EventsCollection
	eg.Go(func() error {
		coll, err := i.ListBirthsFor(ctx, month, day)
		if err != nil {
			return err
		}
		deathEvents = coll
		return nil
	})

	var holidays events.EventsCollection
	eg.Go(func() error {
		coll, err := i.ListBirthsFor(ctx, month, day)
		if err != nil {
			return err
		}
		birthEvents = coll
		return nil
	})
	err := eg.Wait()
	if err != nil {
		return events.EventsCollection{}, fmt.Errorf("failed to list all events for %d-%d: %w", month, day, err)
	}
	aggregatedEvents := make([]events.Event, 0)
	slices.Concat(aggregatedEvents, historicalEvents.Collection, birthEvents.Collection, deathEvents.Collection, holidays.Collection)
	return events.EventsCollection{
		Month:      month,
		Day:        day,
		Collection: aggregatedEvents,
	}, nil
}

// ListBirthsFor implements Interface.
func (i *Impl) ListBirthsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	return i.doList(ctx, events.Birth, month, day)
}

// ListDeathsFor implements Interface.
func (i *Impl) ListDeathsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	return i.doList(ctx, events.Death, month, day)
}

// ListHistoricalEventsFor implements Interface.
func (i *Impl) ListHistoricalEventsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	return i.doList(ctx, events.Historical, month, day)
}

// ListHolidaysFor implements Interface.
func (i *Impl) ListHolidaysFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	return i.doList(ctx, events.Holiday, month, day)
}

func (i *Impl) doList(ctx context.Context, typing events.Type, month, day int64) (events.EventsCollection, error) {
	connRes, ok := i.connectionPool.Get().(*connectionResult)
	if !ok {
		i.logger.Panicf("unexpected type while retrieving connection from pool")
	}
	if connRes.err != nil {
		return events.EventsCollection{},
			fmt.Errorf(`something wrong happened while retrieving connection from pool for upserting events %w`,
				connRes.err)
	}
	conn := connRes.conn
	// NewStargateClientWithConn implementation never returns an error
	cassandraClient, _ := client.NewStargateClientWithConn(conn)
	cql := fmt.Sprintf(`SELECT * FROM %s.events_by_type_day_month WHERE type = ? AND month = ? AND day = ?;`, i.cfg.Keyspace)
	res, err := cassandraClient.ExecuteQueryWithContext(&proto.Query{
		Cql: cql,
		Values: &proto.Values{Values: []*proto.Value{
			{Inner: &proto.Value_String_{String_: typing.String()}},
			{Inner: &proto.Value_Int{Int: month}},
			{Inner: &proto.Value_Int{Int: day}},
		}},
		Parameters: &proto.QueryParameters{
			PageSize: wrapperspb.Int32(math.MaxInt32),
		},
	}, ctx)
	if err != nil {
		return events.EventsCollection{}, DatabaseError{
			Operation: List,
			Type:      events.Holiday,
			Month:     month,
			Day:       day,
			Err:       err,
		}
	}
	coll := events.EventsCollection{
		Type:       events.Holiday,
		Day:        day,
		Month:      month,
		Collection: make([]events.Event, 0),
	}
	for _, r := range res.GetResultSet().GetRows() {
		coll.Collection = append(coll.Collection, events.Event{
			Type:        events.TypeFromString(r.GetValues()[0].GetString_()),
			Day:         r.GetValues()[1].GetInt(),
			Month:       r.GetValues()[2].GetInt(),
			Year:        r.GetValues()[3].GetInt(),
			ID:          r.GetValues()[4].GetString_(),
			Description: r.GetValues()[5].GetString_(),
			Thumbnail: events.Thumbnail{
				Path:   r.GetValues()[7].GetString_(),
				Height: r.GetValues()[6].GetInt(),
				Width:  r.GetValues()[8].GetInt(),
			},
			Title: r.GetValues()[9].GetString_(),
		})
	}
	return coll, nil
}

// UpsertEvents implements Interface.
func (i *Impl) UpsertEvents(ctxc context.Context, coll events.EventsCollection) error {
	connRes, ok := i.connectionPool.Get().(*connectionResult)
	if !ok {
		i.logger.Panicf("unexpected type while retrieving connection from pool")
	}
	if connRes.err != nil {
		return fmt.Errorf("something wrong happened while retrieving connection from pool for upserting events %w", connRes.err)
	}
	conn := connRes.conn
	// NewStargateClientWithConn implementation never returns an error
	cassandraClient, _ := client.NewStargateClientWithConn(conn)
	batchUpsert := i.mapEventsCollectionsToBatchInsert(coll)
	_, err := cassandraClient.ExecuteBatchWithContext(batchUpsert, ctxc)
	if err != nil {
		return DatabaseError{
			Operation: Upsert,
			Type:      coll.Type,
			Month:     coll.Month,
			Day:       coll.Day,
			Err:       err,
		}
	}
	return nil
}

func (i *Impl) mapEventsCollectionsToBatchInsert(coll events.EventsCollection) *proto.Batch {
	batchUpsert := &proto.Batch{
		Type:    proto.Batch_UNLOGGED,
		Queries: make([]*proto.BatchQuery, len(coll.Collection)),
	}
	cql := fmt.Sprintf(`
		INSERT INTO %s.events_by_type_day_month (
			type, day, month, year, id, title ,description , 
			thumbnail_source , thumbnail_width , thumbnail_height )
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`, i.cfg.Keyspace)
	for i := range len(coll.Collection) {
		batchUpsert.Queries[i] = &proto.BatchQuery{
			Cql: cql,
			Values: &proto.Values{
				Values: []*proto.Value{
					{Inner: &proto.Value_String_{String_: coll.Collection[i].Type.String()}},
					{Inner: &proto.Value_Int{Int: coll.Collection[i].Day}},
					{Inner: &proto.Value_Int{Int: coll.Collection[i].Month}},
					{Inner: &proto.Value_Int{Int: coll.Collection[i].Year}},
					{Inner: &proto.Value_String_{String_: coll.Collection[i].ID}},
					{Inner: &proto.Value_String_{String_: coll.Collection[i].Title}},
					{Inner: &proto.Value_String_{String_: coll.Collection[i].Description}},
					{Inner: &proto.Value_String_{String_: coll.Collection[i].Thumbnail.Path}},
					{Inner: &proto.Value_Int{Int: coll.Collection[i].Thumbnail.Width}},
					{Inner: &proto.Value_Int{Int: coll.Collection[i].Thumbnail.Height}},
				},
			},
		}
	}
	return batchUpsert
}
