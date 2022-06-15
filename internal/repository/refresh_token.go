package repository

import (
	"context"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type token struct{}

func NewTokenRepo() TokenRepository {
	return &token{}
}

func (token) Create(ctx context.Context, user *model.User) (*model.RefreshToken, error) {
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

func (token) Find(ctx context.Context, token string) (*model.RefreshToken, error) {
	var refreshToken model.RefreshToken

	err := app.DB().Where(&model.RefreshToken{
		Revoked: false,
		Token:   token,
	}).First(&refreshToken).Error
	err = checkError(err)

	return &refreshToken, err
}

func (token) Revoke(ctx context.Context, token string) error {
	return app.DB().Where(&model.RefreshToken{
		Token: token,
	}).Update("revoked", 1).Error
}
