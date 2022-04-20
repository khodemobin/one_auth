package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/helper"
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
	err = checkError(err)
	if err != nil {
		return nil, err
	}

	jsonToken, err := helper.ToJson(refreshToken)
	if err != nil {
		return nil, err
	}
	err = app.Cache().Set(fmt.Sprintf("refresh_token_%s", refreshToken.Token), jsonToken, 0)

	return refreshToken, err
}

func (token) Find(ctx context.Context, token string) (*model.RefreshToken, error) {
	cache, err := app.Cache().Get(fmt.Sprintf("refresh_token_%s", token), func() (*string, error) {
		var t *model.RefreshToken

		err := app.DB().Where(&model.RefreshToken{
			Revoked: false,
			Token:   token,
		}).First(&t).Error
		err = checkError(err)
		if err != nil {
			return nil, err
		}

		json, err := helper.ToJson(t)
		if err != nil {
			return nil, err
		}

		return &json, nil
	})
	if err != nil {
		return nil, err
	}

	var refreshToken model.RefreshToken
	err = json.Unmarshal([]byte(*cache), &refreshToken)
	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (token) Revoke(ctx context.Context, token *model.RefreshToken) error {
	err := app.Cache().Delete(fmt.Sprintf("refresh_token_%s", token.Token))
	if err != nil {
		return err
	}

	return app.DB().Model(token).Update("revoked", 1).Error
}
