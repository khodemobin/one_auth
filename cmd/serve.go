package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/khodemobin/pilo/auth/app"

	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/internal/server"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
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
	repository := repository.NewRepository()
	service := service.NewService(repository)

	// start server
	restServer := server.New(service, helper.IsLocal())
	go func() {
		if err := restServer.Start(helper.IsLocal(), app.Config().App.Port); err != nil {
			msg := fmt.Sprintf("error happen while serving: %v", err)
			app.Log().Error(errors.New(msg))
			log.Println(msg)
		}
	}()

	// start queue
	app.Queue().Start()

	// wait for close signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("Received an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		fmt.Println("Rest server doesn't shutdown in 10 seconds")
	}
}
