package app

import (
	redisDriver "github.com/go-redis/redis/v8"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/broker"
	"github.com/khodemobin/pilo/auth/pkg/broker/rabbit"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/cache/redis"
	"github.com/khodemobin/pilo/auth/pkg/db"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/logger/sentry"
	"github.com/khodemobin/pilo/auth/pkg/logger/syslog"
	"github.com/khodemobin/pilo/auth/pkg/logger/zap"
	redisClient "github.com/khodemobin/pilo/auth/pkg/redis"

	"gorm.io/gorm"
)

type AppContainer struct {
	Cache  cache.Cache
	DB     *gorm.DB
	Redis  *redisDriver.Client
	Log    logger.Logger
	Config *config.Config
	Broker broker.Broker
}

var Container *AppContainer = nil

func New() {
	config := config.New()

	var logger logger.Logger
	if helper.IsLocal() {
		logger = zap.New()
	} else if config.App.Env == "test" {
		logger = syslog.New()
	} else {
		logger = sentry.New(Container.Config)
	}

	broker := rabbit.New(config, logger)
	db := db.New(config, logger).DB
	rc := redisClient.New(config, logger)
	cache := redis.New(rc, logger)

	Container = &AppContainer{
		Config: config,
		Log:    logger,
		Broker: broker,
		DB:     db,
		Cache:  cache,
	}
}

func App() *AppContainer {
	return Container
}

func Cache() cache.Cache {
	return Container.Cache
}

func DB() *gorm.DB {
	return Container.DB
}

func Redis() *redisDriver.Client {
	return Container.Redis
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
