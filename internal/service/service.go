package service

import (
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/messenger"
)

type Service struct {
	LoginService    domain.LoginService
	RegisterService domain.RegisterService
}

func NewService(repo *repository.Repository, logger logger.Logger, cache cache.Cache, msg messenger.Messenger, cfg *config.Config) *Service {
	login := NewLoginService(repo, msg, cfg)
	register := NewRegisterService(repo, msg, cache, cfg)

	return &Service{
		LoginService:    login,
		RegisterService: register,
	}
}
