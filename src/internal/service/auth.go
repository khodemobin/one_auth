package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type login struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) AuthService {
	return &login{
		repo: repo,
	}
}

func (l *login) Login(ctx context.Context, phone, password string, ac *model.Activity) (*Auth, error) {
	user, err := l.repo.UserRepo.FindActive(ctx, "phone", phone)
	if errors.Is(err, app.ErrNotFound) {
		return nil, errors.New("invalid credentials")
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can find user. err : %s", err.Error()))
	}

	if !encrypt.Check(user.Password.String, password) {
		return nil, errors.New("invalid credentials")
	}

	refreshToken, token := l.generateToken(ctx, user)
	l.updateUserStatics(ctx, user, ac)

	return &Auth{
		Token: token,
		RefreshToken: model.RefreshToken{
			Token: refreshToken.Token,
		},
		ExpiresIn: 3600, // 1 hour
		ID:        user.ID,
	}, nil
}

func (l *login) Logout(ctx context.Context, accessToken string, token string, ac *model.Activity) error {
	currentToken, err := l.repo.TokenRepo.Find(ctx, token)

	if errors.Is(err, app.ErrNotFound) {
		return nil
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can not find refresh token from db. err : %s", err.Error()))
	}

	var wg sync.WaitGroup
	wg.Add(3)
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
	go func() {
		if err := l.repo.ActivityRepo.Create(ac); err != nil {
			panic(fmt.Sprintf("internal error, can not create activity log. err : %s", err.Error()))
		}
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func (l *login) generateToken(ctx context.Context, user *model.User) (*model.RefreshToken, string) {
	refreshToken, err := l.repo.TokenRepo.Create(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	token, err := encrypt.GenerateAccessToken(user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	return refreshToken, token
}

func (l *login) updateUserStatics(ctx context.Context, user *model.User, ac *model.Activity) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		if err := l.repo.UserRepo.UpdateLastSeen(ctx, user); err != nil {
			panic(fmt.Sprintf("internal error, can not update user last seen err : %s", err.Error()))
		}
		wg.Done()
	}()

	go func() {
		if err := l.repo.ActivityRepo.Create(ac); err != nil {
			panic(fmt.Sprintf("internal error, can not create activity log. err : %s", err.Error()))
		}
		wg.Done()
	}()

	wg.Wait()
}
