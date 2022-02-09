package cmd

import (
	"database/sql"
	"github.com/khodemobin/pilo/auth/pkg/db"

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
			sql, _ := sql.Open("mysql", db.Dsn(config))
			dir := "migrations"
			err := goose.SetDialect("mysql")
			if err != nil {
				logger.Fatal(err)
			}
			if err := goose.Run(args[0], sql, dir, args[1:]...); err != nil {
				logger.Fatal(err)
			}
		},
	}

	return cmdMigrate
}
