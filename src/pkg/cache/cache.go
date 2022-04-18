package cache

import "time"

type Cache interface {
	Get(key string, defaultValue func() (*string, error)) (*string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
	Pull(key string, defaultValue func() (*string, error)) (*string, error)
	Close()
}
