package cache

type Cache interface {
	Get(key string, defaultValue func() (*string, error)) (*string, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	Pull(key string, defaultValue func() (*string, error)) (*string, error)
	Close()
}
