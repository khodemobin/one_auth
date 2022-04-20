package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type refresh struct {
	repo *repository.Repository
}

func NewRefreshTokenService(repo *repository.Repository) RefreshTokenService {
	return &refresh{
		repo: repo,
	}
}

func (r *refresh) RefreshToken(ctx context.Context, tokenString string, ac *model.Activity) (*Auth, error) {
	currentToken, err := r.checkRefreshTokenValid(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	user, err := r.repo.UserRepo.FindUserByID(ctx, currentToken.UserID, 1)
	if errors.Is(err, app.ErrNotFound) {
		return nil, errors.New("invalid refresh token")
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	refreshToken, token := r.generateToken(ctx, user)
	r.deleteRefreshToken(ctx, currentToken)

	if err := r.repo.ActivityRepos.CreateActivity(ac); err != nil {
		panic(fmt.Sprintf("internal error, can not create activity log. err : %s", err.Error()))
	}

	return &Auth{
		Token: token,
		RefreshToken: model.RefreshToken{
			Token: refreshToken.Token,
		},
		ExpiresIn: 3600, // 1 hour
		ID:        user.UUID,
	}, nil
}

func (r *refresh) checkRefreshTokenValid(ctx context.Context, tokenString string) (*model.RefreshToken, error) {
	currentToken, err := r.repo.TokenRepo.FindToken(ctx, tokenString)
	if errors.Is(err, app.ErrNotFound) {
		return nil, errors.New("invalid refresh token")
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can not check token exists in db. err : %s", err.Error()))
	}

	if currentToken.CreatedAt.Add(time.Hour*720).Unix() <= time.Now().Unix() {
		r.deleteRefreshToken(ctx, currentToken)
		return nil, errors.New("invalid refresh token")
	}

	return currentToken, nil
}

func (r *refresh) generateToken(ctx context.Context, user *model.User) (*model.RefreshToken, string) {
	refreshToken, err := r.repo.TokenRepo.CreateToken(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	token, err := encrypt.GenerateAccessToken(user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	return refreshToken, token
}

func (r *refresh) deleteRefreshToken(ctx context.Context, token *model.RefreshToken) {
	if err := r.repo.TokenRepo.RevokeToken(ctx, token); err != nil {
		panic(fmt.Sprintf("internal error, can not delete refresh token from db. err : %s", err.Error()))
	}
}
