package redis

import (
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

var client *redis.Client

func New(cfg *config.Config, logger logger.Logger) *redis.Client {
	db, err := strconv.Atoi(cfg.Redis.Database)
	if err != nil {
		logger.Fatal(err)
	}

	poolSize, err := strconv.Atoi(cfg.Redis.PoolSize)
	if err != nil {
		logger.Fatal(err)
	}

	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       db,
		PoolSize: poolSize,
	})

	return client
}

func Get() *redis.Client {
	return client
}
