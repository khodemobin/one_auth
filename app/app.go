package app

import (
	"github.com/khodemobin/pilo/auth/config"
	"github.com/khodemobin/pilo/auth/pkg/broker"
	"github.com/khodemobin/pilo/auth/pkg/db/mysql"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/logger/sentry"
	"github.com/khodemobin/pilo/auth/pkg/logger/syslog"
	"github.com/khodemobin/pilo/auth/pkg/logger/zap"
	"github.com/khodemobin/pilo/auth/pkg/utils"

	"gorm.io/gorm"
)

type AppContainer struct {
	DB     *gorm.DB
	Log    logger.Logger
	Config *config.Config
	Broker broker.Broker
}

var Container *AppContainer = nil

func New() {
	cfg := config.New()

	var log logger.Logger
	if utils.IsLocal() {
		log = zap.New()
	} else if cfg.App.Env == "test" {
		log = syslog.New()
	} else {
		log = sentry.New(Container.Config)
	}

	rabbit := broker.NewRabbitMQ(cfg)
	db, err := mysql.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	Container = &AppContainer{
		Config: cfg,
		Log:    log,
		Broker: rabbit,
		DB:     db,
	}
}

func DB() *gorm.DB {
	return Container.DB
}

func Log() logger.Logger {
	return Container.Log
}

func Config() *config.Config {
	return Container.Config
}

func Broker() broker.Broker {
	return Container.Broker
}
