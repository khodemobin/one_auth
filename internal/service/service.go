package service

import (
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/messager"
)

type Service struct {
	AuthService domain.AuthService
}

func NewService(repo *repository.Repository, logger logger.Logger, msg messager.Messenger, cfg *config.Config) *Service {
	auth := NewAuthService(repo, logger, cfg)

	return &Service{
		AuthService: auth,
	}
}
