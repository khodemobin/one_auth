package cmd

import (
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/spf13/cobra"
)

func SeedCommand() *cobra.Command {
	cmdSeed := &cobra.Command{
		Use:   "seed",
		Short: "Insert fake data to db",
		Run: func(cmd *cobra.Command, args []string) {
			pass, _ := encrypt.Hash("123456")
			user, _ := domain.User{}.SeedUser()
			user.Phone = "09384642495"
			user.Phone = pass
			app.DB().Create(user)
		},
	}

	return cmdSeed
}
