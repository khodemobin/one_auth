package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/internal/server"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/logger/sentry"
	"github.com/khodemobin/pilo/auth/pkg/logger/zap"
	"github.com/khodemobin/pilo/auth/pkg/messager/rabbit"
	"github.com/khodemobin/pilo/auth/pkg/mysql"
	"github.com/khodemobin/pilo/auth/pkg/redis"
)

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
	db := mysql.New(config, logger)
	redis := redis.New(config, logger)

	cache := cache.New(redis, logger)

	defer db.Close()
	defer cache.Close()

	repository := repository.NewRepository(db.DB, cache)
	service := service.NewService(repository, logger, msg)

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
