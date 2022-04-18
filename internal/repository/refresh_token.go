package repository

import (
	"context"
	"errors"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"gorm.io/gorm"
)

type token struct{}

func NewTokenRepo() TokenRepository {
	return &token{}
}

func (token) CreateToken(ctx context.Context, user *model.User) (*model.RefreshToken, error) {
	token, err := encrypt.SecureToken()
	if err != nil {
		return nil, err
	}

	refreshToken := &model.RefreshToken{
		Token:   token,
		UserID:  user.ID,
		Revoked: false,
	}

	err = app.DB().Create(refreshToken).Error
	return refreshToken, err
}

func (token) FindToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	var t *model.RefreshToken

	err := app.DB().Where(&model.RefreshToken{
		Revoked: false,
		Token:   token,
	}).First(&t).Error

	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return t, nil
	}

	return nil, err
}

func (token) RevokeToken(ctx context.Context, token *model.RefreshToken) error {
	return nil
}
