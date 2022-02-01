package service

import (
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/messager"
)

type Service struct {
	AuthService domain.AuthService
}

func NewService(repo *repository.Repository, logger logger.Logger, msg messager.Messager) *Service {
	auth := NewAuthService(logger)

	return &Service{
		AuthService: auth,
	}
}
