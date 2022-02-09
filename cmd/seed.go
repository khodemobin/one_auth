package cmd

import (
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/db"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/logger/sentry"
	"github.com/khodemobin/pilo/auth/pkg/logger/zap"
	"github.com/spf13/cobra"
)

func SeedCommand() *cobra.Command {
	config := config.New()
	var logger logger.Logger
	if helper.IsLocal(config) {
		logger = zap.New()
	} else {
		logger = sentry.New(config)
	}

	cmdSeed := &cobra.Command{
		Use:   "seed",
		Short: "Insert fake data to db",
		Run: func(cmd *cobra.Command, args []string) {
			db := db.New(config, logger)
			pass, _ := encrypt.Hash("123456")
			user := &domain.User{
				Phone:              "09384642495",
				Password:           &pass,
				ConfirmationToken:  nil,
				ConfirmationSentAt: nil,
				ConfirmedAt:        nil,
				RecoveryToken:      nil,
				RecoverySentAt:     nil,
				PhoneChangeToken:   nil,
				PhoneChange:        nil,
				PhoneChangeSentAt:  nil,
				Role:               nil,
				Status:             domain.USER_STATUS_ACTIVE,
				IsSuperAdmin:       false,
				LastSignInAt:       nil,
				Tokens:             nil,
			}

			db.DB.Create(user)
		},
	}

	return cmdSeed
}
