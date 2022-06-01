package test_mock

import (
	"os"
	"testing"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/pkg/cache/redis"
	"github.com/khodemobin/pilo/auth/pkg/logger/syslog"
)

func NewMock(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	app.New()
	logger := syslog.New()
	redis := redis.NewTest(t, logger)
	app.Container.Cache = redis
}
