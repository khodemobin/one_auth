package test_mock

import (
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/cache/redis"
	"github.com/khodemobin/pilo/auth/pkg/db"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/logger/syslog"
	"gorm.io/gorm"
	"testing"
)

func NewMock(t *testing.T) (*gorm.DB, cache.Cache, logger.Logger) {
	logger := syslog.New()
	cfg := config.GetConfig()
	db := db.New(cfg, logger)
	redis := redis.NewTest(t, logger)

	return db.DB, redis, logger
}
