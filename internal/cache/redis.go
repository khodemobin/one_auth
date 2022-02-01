package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type redisCache struct {
	rc     *redis.Client
	logger logger.Logger
}

var ctx = context.Background()

func New(rc *redis.Client, logger logger.Logger) Cache {
	return &redisCache{
		rc:     rc,
		logger: logger,
	}
}

func Close(cache *redis.Client) error {
	return cache.Close()
}

func (r *redisCache) Get(key string, defaultValue func() (*string, error)) (*string, error) {
	value, err := r.rc.Get(ctx, helper.ToMD5(key)).Result()
	if err != nil {
		if defaultValue == nil {
			return nil, nil
		}

		v, err := defaultValue()
		if err != nil {
			return nil, err
		}

		err = r.Set(key, *v)

		if err != nil {
			return nil, err
		}

		return defaultValue()
	}

	return &value, err
}

func (r *redisCache) Set(key string, value interface{}) error {
	return r.rc.Set(ctx, helper.ToMD5(key), value, 0).Err()
}

func (r *redisCache) Delete(key string) error {
	return r.rc.Del(ctx, helper.ToMD5(key)).Err()
}

func (r *redisCache) Pull(key string, defaultValue func() (*string, error)) (*string, error) {
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

func (r *redisCache) Close() {
	err := r.rc.Close()
	if err != nil {
		r.logger.Fatal(err)
	}
}
