package cmd

import (
	"errors"
	"fmt"
	"github.com/khodemobin/pilo/auth/pkg/cache/redis"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/internal/server"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/db"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/logger/sentry"
	"github.com/khodemobin/pilo/auth/pkg/logger/zap"
	"github.com/khodemobin/pilo/auth/pkg/messager/rabbit"
	"github.com/spf13/cobra"
)

func ServeCommand() *cobra.Command {
	cmdServe := &cobra.Command{
		Use:   "serve",
		Short: "Serve application",
		Run: func(cmd *cobra.Command, args []string) {
			Execute()
		},
	}
	return cmdServe
}

func Execute() {
	// init main components
	config := config.New()

	var logger logger.Logger
	if helper.IsLocal(config) {
		logger = zap.New()
	} else {
		logger = sentry.New(config)
	}

	msg := rabbit.New(config, logger)
	db := db.New(config, logger)
	redis := redis.New(config, logger)

	defer db.Close()
	defer redis.Close()

	repository := repository.NewRepository(db.DB, redis)
	service := service.NewService(repository, logger, msg, config)

	// start server
	restServer := server.New(service, helper.IsLocal(config), logger)
	go func() {
		if err := restServer.Start(helper.IsLocal(config), config.App.Port); err != nil {
			msg := fmt.Sprintf("error happen while serving: %v", err)
			logger.Error(errors.New(msg))
			log.Println(msg)
		}
	}()

	// wait for close signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("Received an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		fmt.Println("Rest server doesn't shutdown in 10 seconds")
	}
}
