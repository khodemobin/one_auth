package service

import (
	"context"
	"errors"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"strconv"
)

type auth struct {
	repo   *repository.Repository
	logger logger.Logger
	cfg    *config.Config
}

func NewAuthService(repo *repository.Repository, logger logger.Logger, cfg *config.Config) domain.AuthService {
	return &auth{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

func (a *auth) Login(ctx context.Context, phone, password string) (*domain.Auth, error) {
	user, err := a.repo.UserRepo.FindUserByPhone(ctx, phone)
	if err != nil {
		panic(err)
	}

	if user.ID == 0 {
		return nil, errors.New("invalid credentials")
	}

	if !encrypt.Check(*user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	ttl, err := strconv.Atoi(a.cfg.App.JwtTTL)
	if err != nil {
		// TODO do not use panic. handle custom error 500
		panic("internal error, can not convert jwt time to int type")
	}

	token, err := a.repo.TokenRepo.Create(ctx, ttl, user)
	if err != nil {
		panic("internal error, can not create token")
	}

	err = a.repo.UserRepo.UpdateUserLastSeen(ctx, user)
	if err != nil {
		panic("internal error, can not create token")
	}

	// TODO add events log and back and security

	return &domain.Auth{
		Token:     token.Token,
		ExpiresIn: ttl,
		UserID:    user.ID,
	}, nil
}

func (a *auth) Check(ctx context.Context) (*domain.Auth, error) {
	return nil, nil
}
