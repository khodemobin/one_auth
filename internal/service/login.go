package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
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

	ttl, err := strconv.Atoi(app.Config().App.JwtTTL)
	if err != nil {
		panic(err)
	}

	token, err := l.repo.TokenRepo.CreateToken(ctx, ttl, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	if err := l.repo.ActivityRepos.CreateActivity(ac); err != nil {
		panic(err)
	}

	return &Auth{
		Token:     token.Token,
		ExpiresIn: ttl,
		ID:        user.UUID,
	}, nil
}

func (*login) Logout(ctx context.Context, token string, ac *model.Activity) error {
	return nil
}
