package redis

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/helper"

	"github.com/go-redis/redis/v8"

	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type client struct {
	rc     *redis.Client
	ctx    context.Context
	logger logger.Logger
}

func New(rc *redis.Client, logger logger.Logger) cache.Cache {
	return &client{
		rc:     rc,
		logger: logger,
		ctx:    context.Background(),
	}
}

func NewTest(t *testing.T, logger logger.Logger) cache.Cache {
	s := miniredis.RunT(t)
	r := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	return &client{
		rc:     r,
		logger: logger,
		ctx:    context.Background(),
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
