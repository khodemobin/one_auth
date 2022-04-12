package service

import (
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
)

type Service struct {
	LoginService    domain.LoginService
	RegisterService domain.RegisterService
}

func NewService(repo *repository.Repository) *Service {
	login := NewLoginService(repo)
	register := NewRegisterService(repo)

	return &Service{
		LoginService:    login,
		RegisterService: register,
	}
}
