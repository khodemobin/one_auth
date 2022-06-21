package cmd

import (
	"database/sql"
	"github.com/khodemobin/pilo/auth/pkg/db/mysql"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

func MigrateCommand() *cobra.Command {
	cmdMigrate := &cobra.Command{
		Use:   "migrate [ up & down & create]",
		Short: "Migrate database [ up & down & create]",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sql, _ := sql.Open("mysql", mysql.Dsn(app.Config()))
			dir := "migrations"
			err := goose.SetDialect("mysql")
			if err != nil {
				app.Log().Fatal(err)
			}
			if err := goose.Run(args[0], sql, dir, args[1:]...); err != nil {
				app.Log().Fatal(err)
			}
		},
	}

	return cmdMigrate
}
