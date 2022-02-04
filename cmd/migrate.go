package cmd

import (
	"database/sql"
	"fmt"

	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/logger/sentry"
	"github.com/khodemobin/pilo/auth/pkg/logger/zap"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

func MigrateCommand() *cobra.Command {
	config := config.New()
	var logger logger.Logger
	if helper.IsLocal(config) {
		logger = zap.New()
	} else {
		logger = sentry.New(config)
	}

	cmdMigrate := &cobra.Command{
		Use:   "migrate [ up & down & create]",
		Short: "Migrate database [ up & down & create]",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := sql.Open("mysql", dsn(config))
			dir := "migrations"
			goose.SetDialect("mysql")
			if err := goose.Run(args[0], db, dir, args[1:]...); err != nil {
				logger.Fatal(err)
			}
		},
	}

	return cmdMigrate
}

func dsn(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
}
