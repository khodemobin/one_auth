package service

import (
	"context"
	"fmt"
	"github.com/khodemobin/pilo/auth/internal/http/request"
	"github.com/khodemobin/pilo/auth/pkg/utils"
	encrypt2 "github.com/khodemobin/pilo/auth/pkg/utils/encrypt"
	"sync"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
)

type login struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) AuthService {
	return &login{
		repo: repo,
	}
}

func (l *login) Login(ctx context.Context, req request.LoginRequest) (*Auth, error) {
	user, err := l.repo.UserRepo.FindActive(ctx, "phone", req.Phone)
	if errors.Is(err, app.ErrNotFound) {
		return nil, errors.New("invalid credentials")
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can find user. err : %s", err.Error()))
	}

	if !encrypt2.Check(user.Password.String, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	refreshToken, token := l.generateToken(ctx, user)
	if err := l.repo.UserRepo.UpdateLastSeen(ctx, user); err != nil {
		panic(fmt.Sprintf("internal error, can not update user last seen err : %s", err.Error()))
	}

	return &Auth{
		Token: token,
		RefreshToken: model.RefreshToken{
			Token: refreshToken.Token,
		},
		ExpiresIn: 3600, // 1 hour
		ID:        user.ID,
	}, nil
}

func (l *login) Logout(ctx context.Context, accessToken string, token string) error {
	currentToken, err := l.repo.TokenRepo.Find(ctx, token)

	if errors.Is(err, app.ErrNotFound) {
		return nil
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can not find refresh token from db. err : %s", err.Error()))
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if err := l.repo.TokenRepo.Revoke(ctx, currentToken.Token); err != nil {
			panic(fmt.Sprintf("internal error, can not delete refresh token from db. err : %s", err.Error()))
		}
		wg.Done()
	}()

	go func() {
		if err := l.repo.BlackListRepo.Create(accessToken); err != nil {
			panic(fmt.Sprintf("internal error, can not add access token to blacklist. err : %s", err.Error()))
		}
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func (l *login) Refresh(ctx context.Context, tokenString string) (*Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (l *login) generateToken(ctx context.Context, user *model.User) (*model.RefreshToken, string) {
	refreshToken, err := l.repo.TokenRepo.Create(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	token, err := utils.GenerateAccessToken(user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	return refreshToken, token
}
