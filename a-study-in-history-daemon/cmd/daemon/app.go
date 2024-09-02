package main

import (
	"context"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/facade/importer"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/cassandra"
	"github.com/aaegamysta/a-study-in-history/daemon/internal/infrastructure/wikipedia"
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type App struct {
	logger *zap.SugaredLogger
}

func CreateApp(_ context.Context) App {
	zapConfig := zap.NewDevelopmentConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		log.Panicf("failed to create logger while creatinga application %v", err)
	}
	sugaredLogger := logger.Sugar()
	f, err := os.Open("./configs/daemon.development.yaml")
	if err != nil {
		log.Println(os.Getwd())
		log.Panicf("failed to open config file while creating application %v", err)
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Panicf("failed to parse config file while creating application %v", err)
	}
	var cfg = AppConfig{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Panicf("failed to parse config file while creating application %v", err)
	}
	app := App{
		logger: sugaredLogger,
	}
	return app
}

type AppConfig struct {
	Cassandra cassandra.Config `yaml:"cassandra"`
	Wikipedia wikipedia.Config `yaml:"wikipedia"`
	Importer  importer.Config  `yaml:"parser"`
}
