package mysql

import (
	"fmt"
	"github.com/khodemobin/pilo/auth/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(Dsn(cfg)), &gorm.Config{
		Logger: l.Default.LogMode(l.Silent),
	})
}

func Dsn(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local&parseTime=true&multiStatements=true", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
}
