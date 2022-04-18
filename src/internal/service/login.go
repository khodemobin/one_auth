package service

import (
	"context"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type login struct {
	repo *repository.Repository
}

func NewLoginService(repo *repository.Repository) LoginService {
	return &login{
		repo: repo,
	}
}

func (l *login) Login(ctx context.Context, phone, password string, ac *model.Activity) (*Auth, error) {
	user, err := l.repo.UserRepo.FindUserByPhone(ctx, phone, model.USER_STATUS_ACTIVE)
	if err != nil {
		panic(fmt.Sprintf("internal error, can find user. err : %s", err.Error()))
	}

	if user.ID == 0 {
		return nil, errors.New("invalid credentials")
	}

	// TODO active password check
	// if !encrypt.Check(*user.Password, password) {
	// 	return nil, errors.New("invalid credentials")
	// }

	refreshToken, err := l.repo.TokenRepo.CreateToken(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	token, err := encrypt.GenerateAccessToken(user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	if err := l.repo.ActivityRepos.CreateActivity(ac); err != nil {
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

func (*login) Logout(ctx context.Context, token string, ac *model.Activity) error {
	return nil
}

func (l *login) RefreshToken(ctx context.Context, tokenString string, ac *model.Activity) (*Auth, error) {
	token, err := l.repo.TokenRepo.FindToken(ctx, tokenString)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not check token exists in db. err : %s", err.Error()))
	}

	if token.ID == 0 {
		return nil, errors.New("invalid refresh token")
	}

	return nil, nil
}
