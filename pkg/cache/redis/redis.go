package redis

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/helper"

	"github.com/go-redis/redis/v8"
	r "github.com/go-redis/redis/v8"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type client struct {
	rc     *r.Client
	logger logger.Logger
	ctx    context.Context
}

func New(cfg *config.Config, logger logger.Logger) cache.Cache {
	db, err := strconv.Atoi(cfg.Redis.Database)
	if err != nil {
		logger.Fatal(err)
	}

	poolSize, err := strconv.Atoi(cfg.Redis.PoolSize)
	if err != nil {
		logger.Fatal(err)
	}
	r := r.NewClient(&r.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       db,
		PoolSize: poolSize,
	})

	return &client{
		rc:     r,
		logger: logger,
		ctx:    context.Background(),
	}
}

func NewTest(t *testing.T, logger logger.Logger) cache.Cache {
	s := miniredis.RunT(t)
	r := r.NewClient(&r.Options{
		Addr: s.Addr(),
	})
	return &client{
		rc:     r,
		logger: logger,
	}
}

func (r *client) Get(key string, defaultValue func() (*string, error)) (*string, error) {
	value, err := r.rc.Get(r.ctx, helper.ToMD5(key)).Result()
	if err == redis.Nil {
		if defaultValue == nil {
			return nil, nil
		}

		v, err := defaultValue()
		if err != nil {
			return nil, err
		}

		err = r.Set(key, *v, 0)

		if err != nil {
			return nil, err
		}

		return defaultValue()
	}

	return &value, err
}

func (r *client) Set(key string, value interface{}, expiration time.Duration) error {
	return r.rc.Set(r.ctx, helper.ToMD5(key), value, expiration).Err()
}

func (r *client) Delete(key string) error {
	return r.rc.Del(r.ctx, helper.ToMD5(key)).Err()
}

func (r *client) Pull(key string, defaultValue func() (*string, error)) (*string, error) {
	value, err := r.Get(key, defaultValue)
	if err != nil {
		return nil, err
	}

	err = r.Delete(key)

	if err != nil {
		return nil, err
	}

	return value, err
}

func (r *client) Close() {
	err := r.rc.Close()
	if err != nil {
		r.logger.Fatal(err)
	}
}
