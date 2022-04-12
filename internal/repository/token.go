package repository

import (
	"context"
	"time"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type token struct{}

func NewTokenRepo() domain.TokenRepository {
	return &token{}
}

func (token) CreateToken(ctx context.Context, ttl int, user *domain.User) (*domain.Token, error) {
	token, err := encrypt.GenerateAccessToken(user, time.Second*time.Duration(ttl))
	if err != nil {
		return nil, err
	}

	tokenModel := &domain.Token{
		Token:  token,
		UserID: user.ID,
	}

	err = app.DB().Create(&tokenModel).Error
	return tokenModel, err
}

func (token) RevokeToken(ctx context.Context, token *domain.Token) error {
	return nil
}
