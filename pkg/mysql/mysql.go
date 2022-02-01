package mysql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	logger logger.Logger
	DB     *sqlx.DB
}

func New(cfg *config.Config, logger logger.Logger) *Mysql {
	db, err := sqlx.Connect("mysql", dsn(cfg))
	if err != nil {
		logger.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Mysql{
		logger: logger,
		DB:     db,
	}
}

func (m *Mysql) Close() {
	err := m.DB.Close()
	if err != nil {
		m.logger.Fatal(err)
	}
}

func dsn(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
}
