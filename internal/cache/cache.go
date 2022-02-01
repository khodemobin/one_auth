package cache

const (
	DEVICE_BY_IMEI                          = "dbi_"
	USER_DEVICE_BY_DEVICE_ID                = "udbdi_"
	DEVICE_CONFIG_BY_USER_DEVICE_ID         = "dcbudi_"
	DEVICE_CONFIG_CYCLE_BY_DEVICE_CONFIG_ID = "dccbdci_"
)

type Cache interface {
	Get(key string, defaultValue func() (*string, error)) (*string, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	Pull(key string, defaultValue func() (*string, error)) (*string, error)
	Close()
}
