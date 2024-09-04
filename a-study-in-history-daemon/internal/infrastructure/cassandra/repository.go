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

type Operation int64

const (
	Unspecified Operation = iota
	Upsert
	List
)

type Interface interface {
	CreateTablesIfNotExists(ctx context.Context) error
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
	_, err := cassandraClient.ExecuteQuery(&proto.Query{
		Cql: cql,
	})
	if err != nil {
		return err
	}
	return nil
}

// ListBirthsFor implements Interface.
func (i *Impl) ListBirthsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	connRes, ok := i.connectionPool.Get().(*connectionResult)
	if !ok {
		i.logger.Panicf("unexpected type while retrieving connection from pool")
	}
	if connRes.err != nil {
		return events.EventsCollection{}, fmt.Errorf("something wrong happened while retrieving connection from pool for upserting events %w", connRes.err)
	}
	conn := connRes.conn
	// NewStargateClientWithConn implementation never returns an error
	cassandraClient, _ := client.NewStargateClientWithConn(conn)
	cql := fmt.Sprintf(`SELECT * FROM %s.events_by_type_day_month WHERE type = ? AND month = ? AND day = ?;`, i.cfg.Keyspace)
	res, err := cassandraClient.ExecuteQueryWithContext(&proto.Query{
		Cql: cql,
		Values: &proto.Values{Values: []*proto.Value{
			{Inner: &proto.Value_String_{String_: events.Birth.String()}},
			{Inner: &proto.Value_Int{month}},
			{Inner: &proto.Value_Int{day}},
		}},
	}, ctx)
	if err != nil {
		return events.EventsCollection{}, DatabaseError{
			Operation: List,
			Type:      events.Birth,
			Month:     month,
			Day:       day,
			Err:       err,
		}
	}
	coll := events.EventsCollection{
		Type:   events.Birth,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	for _, r := range res.GetResultSet().Rows {
		coll.Events = append(coll.Events, events.Event{
			Type:        events.TypeFromString(r.Values[0].GetString_()),
			Day:         r.Values[1].GetInt(),
			Month:       r.Values[2].GetInt(),
			Year:        r.Values[3].GetInt(),
			ID:          r.Values[4].GetString_(),
			Description: r.Values[5].GetString_(),
			Thumbnail: events.Thumbnail{
				Path:   r.Values[6].GetString_(),
				Height: r.Values[7].GetInt(),
				Width:  r.Values[8].GetInt(),
			},
			Title: r.Values[9].GetString_(),
		})
	}
	return coll, nil
}

// ListDeathsFor implements Interface.
func (i *Impl) ListDeathsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	connRes, ok := i.connectionPool.Get().(*connectionResult)
	if !ok {
		i.logger.Panicf("unexpected type while retrieving connection from pool")
	}
	if connRes.err != nil {
		return events.EventsCollection{}, fmt.Errorf("something wrong happened while retrieving connection from pool for upserting events %w", connRes.err)
	}
	conn := connRes.conn
	// NewStargateClientWithConn implementation never returns an error
	cassandraClient, _ := client.NewStargateClientWithConn(conn)
	cql := fmt.Sprintf(`SELECT * FROM %s.events_by_type_day_month WHERE type = ? AND month = ? AND day = ?;`, i.cfg.Keyspace)
	res, err := cassandraClient.ExecuteQueryWithContext(&proto.Query{
		Cql: cql,
		Values: &proto.Values{Values: []*proto.Value{
			{Inner: &proto.Value_String_{String_: events.Death.String()}},
			{Inner: &proto.Value_Int{month}},
			{Inner: &proto.Value_Int{day}},
		}},
	}, ctx)
	if err != nil {
		return events.EventsCollection{}, DatabaseError{
			Operation: List,
			Type:      events.Death,
			Month:     month,
			Day:       day,
			Err:       err,
		}
	}
	coll := events.EventsCollection{
		Type:   events.Death,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	for _, r := range res.GetResultSet().Rows {
		coll.Events = append(coll.Events, events.Event{
			Type:        events.TypeFromString(r.Values[0].GetString_()),
			Day:         r.Values[1].GetInt(),
			Month:       r.Values[2].GetInt(),
			Year:        r.Values[3].GetInt(),
			ID:          r.Values[4].GetString_(),
			Description: r.Values[5].GetString_(),
			Thumbnail: events.Thumbnail{
				Path:   r.Values[6].GetString_(),
				Height: r.Values[7].GetInt(),
				Width:  r.Values[8].GetInt(),
			},
			Title: r.Values[9].GetString_(),
		})
	}
	return coll, nil
}

// ListHistoricalEventsFor implements Interface.
func (i *Impl) ListHistoricalEventsFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	connRes, ok := i.connectionPool.Get().(*connectionResult)
	if !ok {
		i.logger.Panicf("unexpected type while retrieving connection from pool")
	}
	if connRes.err != nil {
		return events.EventsCollection{}, fmt.Errorf("something wrong happened while retrieving connection from pool for upserting events %w", connRes.err)
	}
	conn := connRes.conn
	// NewStargateClientWithConn implementation never returns an error
	cassandraClient, _ := client.NewStargateClientWithConn(conn)
	cql := fmt.Sprintf(`SELECT * FROM %s.events_by_type_day_month WHERE type = ? AND month = ? AND day = ?;`, i.cfg.Keyspace)
	res, err := cassandraClient.ExecuteQueryWithContext(&proto.Query{
		Cql: cql,
		Values: &proto.Values{Values: []*proto.Value{
			{Inner: &proto.Value_String_{String_: events.Historical.String()}},
			{Inner: &proto.Value_Int{month}},
			{Inner: &proto.Value_Int{day}},
		}},
	}, ctx)
	if err != nil {
		return events.EventsCollection{}, DatabaseError{
			Operation: List,
			Type:      events.Historical,
			Month:     month,
			Day:       day,
			Err:       err,
		}
	}
	coll := events.EventsCollection{
		Type:   events.Historical,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	for _, r := range res.GetResultSet().Rows {
		coll.Events = append(coll.Events, events.Event{
			Type:        events.TypeFromString(r.Values[0].GetString_()),
			Day:         r.Values[1].GetInt(),
			Month:       r.Values[2].GetInt(),
			Year:        r.Values[3].GetInt(),
			ID:          r.Values[4].GetString_(),
			Description: r.Values[5].GetString_(),
			Thumbnail: events.Thumbnail{
				Path:   r.Values[7].GetString_(),
				Height: r.Values[6].GetInt(),
				Width:  r.Values[8].GetInt(),
			},
			Title: r.Values[9].GetString_(),
		})
	}
	return coll, nil
}

// ListHolidaysFor implements Interface.
func (i *Impl) ListHolidaysFor(ctx context.Context, month int64, day int64) (events.EventsCollection, error) {
	connRes, ok := i.connectionPool.Get().(*connectionResult)
	if !ok {
		i.logger.Panicf("unexpected type while retrieving connection from pool")
	}
	if connRes.err != nil {
		return events.EventsCollection{}, fmt.Errorf("something wrong happened while retrieving connection from pool for upserting events %w", connRes.err)
	}
	conn := connRes.conn
	// NewStargateClientWithConn implementation never returns an error
	cassandraClient, _ := client.NewStargateClientWithConn(conn)
	cql := fmt.Sprintf(`SELECT * FROM %s.events_by_type_day_month WHERE type = ? AND month = ? AND day = ?;`, i.cfg.Keyspace)
	res, err := cassandraClient.ExecuteQueryWithContext(&proto.Query{
		Cql: cql,
		Values: &proto.Values{Values: []*proto.Value{
			{Inner: &proto.Value_String_{String_: events.Holiday.String()}},
			{Inner: &proto.Value_Int{month}},
			{Inner: &proto.Value_Int{day}},
		}},
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
		Type:   events.Holiday,
		Day:    day,
		Month:  month,
		Events: make([]events.Event, 0),
	}
	for _, r := range res.GetResultSet().Rows {
		coll.Events = append(coll.Events, events.Event{
			Type:        events.TypeFromString(r.Values[0].GetString_()),
			Day:         r.Values[1].GetInt(),
			Month:       r.Values[2].GetInt(),
			Year:        r.Values[3].GetInt(),
			ID:          r.Values[4].GetString_(),
			Description: r.Values[5].GetString_(),
			Thumbnail: events.Thumbnail{
				Path:   r.Values[6].GetString_(),
				Height: r.Values[7].GetInt(),
				Width:  r.Values[8].GetInt(),
			},
			Title: r.Values[9].GetString_(),
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
		Queries: make([]*proto.BatchQuery, len(coll.Events)),
	}
	cql := fmt.Sprintf(`
		INSERT INTO %s.events_by_type_day_month (
			type, day, month, year, id, title ,description , 
			thumbnail_source , thumbnail_width , thumbnail_height )
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`, i.cfg.Keyspace)
	for i := 0; i < len(coll.Events); i++ {
		batchUpsert.Queries[i] = &proto.BatchQuery{
			Cql: cql,
			Values: &proto.Values{
				Values: []*proto.Value{
					&proto.Value{Inner: &proto.Value_String_{coll.Events[i].Type.String()}},
					&proto.Value{Inner: &proto.Value_Int{coll.Events[i].Day}},
					&proto.Value{Inner: &proto.Value_Int{coll.Events[i].Month}},
					&proto.Value{Inner: &proto.Value_Int{coll.Events[i].Year}},
					&proto.Value{Inner: &proto.Value_String_{coll.Events[i].ID}},
					&proto.Value{Inner: &proto.Value_String_{coll.Events[i].Title}},
					&proto.Value{Inner: &proto.Value_String_{coll.Events[i].Description}},
					&proto.Value{Inner: &proto.Value_String_{coll.Events[i].Thumbnail.Path}},
					&proto.Value{Inner: &proto.Value_Int{coll.Events[i].Thumbnail.Width}},
					&proto.Value{Inner: &proto.Value_Int{coll.Events[i].Thumbnail.Height}},
				},
			},
		}
	}
	return batchUpsert
}
