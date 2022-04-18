package sentry

import (
	l "log"

	"github.com/getsentry/sentry-go"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type log struct{}

func New(cfg *config.Config) logger.Logger {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: cfg.Sentry.Dsn,
	})
	if err != nil {
		l.Fatalf("sentry.Init: %s", err)
	}

	return &log{}
}

func (l *log) Error(err error) {
	if err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelFatal)
			sentry.CaptureException(err)
		})
	}
}

func (l *log) Fatal(err error) {
	if err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelFatal)
			sentry.CaptureException(err)
		})
		l.Fatal(err)
	}
}

func (l *log) Warn(msg string) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelWarning)
		sentry.CaptureMessage(msg)
	})
}

func (l *log) Info(msg string) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelInfo)
		sentry.CaptureMessage(msg)
	})
}
