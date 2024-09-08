package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"

	"github.com/aaegamysta/a-study-in-history/daemon/internal/facade/collector"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/facade/importer"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/facade/resync"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/cassandra"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/wikipedia"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/server"
	"github.com/aaegamysta/a-study-in-history/spec/gen"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

type App struct {
	logger          *zap.SugaredLogger
	WikipediaClient wikipedia.Interface
	CassandraClient cassandra.Interface
	importer        importer.Interface
	server          gen.AStudyInHistoryServer
	gRPCServer      *grpc.Server
}

func CreateApp(ctx context.Context) App {
	zapConfig := zap.NewDevelopmentConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		log.Panicf("failed to create logger while creatinga application %v", err)
	}
	sugaredLogger := logger.Sugar()
	f, err := os.Open("./configs/daemon.development.yaml")
	if err != nil {
		sugaredLogger.Panicf("failed to open config file while creating application %v", err)
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		sugaredLogger.Panicf("failed to parse config file while creating application %v", err)
	}
	cfg := AppConfig{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		sugaredLogger.Panicf("failed to parse config file while creating application %v", err)
	}
	wikipediaClient := wikipedia.New(ctx, cfg.Wikipedia, sugaredLogger)
	cassandraClient := cassandra.New(ctx, cfg.Cassandra, sugaredLogger)
	importer := importer.New(ctx, sugaredLogger, cfg.Importer, cassandraClient, wikipediaClient)
	collector := collector.New(ctx, sugaredLogger, cassandraClient)
	resynchronizer := resync.New(ctx, sugaredLogger, wikipediaClient, cassandraClient)
	server := server.New(ctx, sugaredLogger, importer, resynchronizer, collector)
	app := App{
		logger:          sugaredLogger,
		WikipediaClient: wikipediaClient,
		CassandraClient: cassandraClient,
		importer:        importer,
		server:          server,
	}
	return app
}

type AppConfig struct {
	Cassandra cassandra.Config `yaml:"cassandra"`
	Wikipedia wikipedia.Config `yaml:"wikipedia"`
	Importer  importer.Config  `yaml:"parser"`
}

func (app *App) Run(ctx context.Context) {
	// app.importer.Import(ctx)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		app.logger.Panicf("failed to create listener %v", err)
	}
	gRPCServer := grpc.NewServer()
	app.gRPCServer = gRPCServer
	go func() {
		gen.RegisterAStudyInHistoryServer(gRPCServer, app.server)
		reflection.Register(gRPCServer)
		app.logger.Infof("a study in history server started on port %d", 50051)
		err = gRPCServer.Serve(listener)
		if err != nil {
			app.logger.Panicf("failed to start gRPC server %v", err)
		}
	}()
}

func (app *App) Shutdown() {
	app.logger.Infof("gracefully shutting down server")
	app.gRPCServer.GracefulStop()
	app.logger.Infof("application and server shut down")
}
