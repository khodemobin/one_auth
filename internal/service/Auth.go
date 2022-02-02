package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type auth struct {
	repo   *repository.Repository
	logger *logger.Logger
	cfg    *config.Config
}

func NewAuthService(repo *repository.Repository, logger *logger.Logger, cfg *config.Config) domain.AuthService {
	return &auth{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

func (a *auth) Login(ctx context.Context, phone, password string) (*domain.Auth, error) {
	user, err := a.repo.UserRepo.FindUserByPhone(phone)
	if err != nil {
		return nil, &domain.AuthError{}
	}

	if !encrypt.Check(password, *user.Password) {
		return nil, &domain.AuthError{}
	}

	ttl, err := strconv.Atoi(a.cfg.App.JwtTTL)
	if err != nil {
		return nil, errors.New("internal error, can not convert jwt time to int type")
	}

	token, err := encrypt.GenerateAccessToken(user, time.Second*time.Duration(ttl), a.cfg.App.JwtSecret)
	if err != nil {
		return nil, errors.New("internal error, can not create tok")
	}

	// TODO add events log and back and security

	return &domain.Auth{
		Token: token,
	}, nil
}

func (a *auth) Check(ctx context.Context) (*domain.Auth, error) {
	return nil, nil
}
