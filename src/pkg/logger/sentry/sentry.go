package sentry

import (
	sysLog "log"

	"github.com/getsentry/sentry-go"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type log struct{}

func New[T string | error](cfg *config.Config) logger.Logger {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: cfg.Sentry.Dsn,
	})
	if err != nil {
		sysLog.Fatalf("sentry.Init: %s", err)
	}

	return &log{}
}

func (l *log) Error(msg logger.ErrorType) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelFatal)
		sentry.CaptureException(logger.GetError(msg))
	})
}

func (l *log) Fatal(msg logger.ErrorType) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelFatal)
		sentry.CaptureException(logger.GetError(msg))
	})
	sysLog.Fatal(msg)
}

func (l *log) Warn(msg logger.ErrorType) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelWarning)
		sentry.CaptureException(logger.GetError(msg))
	})
}

func (l *log) Info(msg logger.ErrorType) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelInfo)
		sentry.CaptureException(logger.GetError(msg))
	})
}
