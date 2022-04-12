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
	"github.com/khodemobin/pilo/auth/pkg/logger/zap"
	"github.com/khodemobin/pilo/auth/pkg/queue"
	redisClient "github.com/khodemobin/pilo/auth/pkg/redis"

	"gorm.io/gorm"
)

type AppContainer struct {
	Cache  cache.Cache
	DB     *gorm.DB
	Queue  *queue.Queue
	Redis  *redisDriver.Client
	Log    logger.Logger
	Config *config.Config
	Broker broker.Broker
}

var app *AppContainer = nil

func New() {
	config := config.New()

	var logger logger.Logger
	if helper.IsLocal() {
		logger = zap.New()
	} else {
		logger = sentry.New(app.Config)
	}

	broker := rabbit.New(config, logger)
	db := db.New(config, logger).DB
	rc := redisClient.New(config, logger)
	cache := redis.New(rc, logger)
	queue := queue.New(rc, logger)

	app = &AppContainer{
		Config: config,
		Log:    logger,
		Broker: broker,
		DB:     db,
		Cache:  cache,
		Queue:  queue,
	}
}

func App() *AppContainer {
	return app
}

func Cache() cache.Cache {
	return app.Cache
}

func DB() *gorm.DB {
	return app.DB
}

func Queue() *queue.Queue {
	return app.Queue
}

func Redis() *redisDriver.Client {
	return app.Redis
}

func Log() logger.Logger {
	return app.Log
}

func Config() *config.Config {
	return app.Config
}

func Broker() broker.Broker {
	return app.Broker
}
