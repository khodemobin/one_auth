package service

import (
	"context"

	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type auth struct {
	// repo   *repository.Repository
	logger logger.Logger
}

func NewAuthService(logger logger.Logger) domain.AuthService {
	return &auth{
		logger: logger,
	}
}

func (f *auth) Login(ctx context.Context) (string, error) {
	return "", nil
}

func (f *auth) Check(ctx context.Context) (string, error) {
	return "", nil
}
