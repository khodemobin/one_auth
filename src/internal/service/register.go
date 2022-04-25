package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type register struct {
	repo *repository.Repository
}

func NewRegisterService(repo *repository.Repository) RegisterService {
	return &register{
		repo: repo,
	}
}

func (r *register) Request(ctx context.Context, phone string, ac *model.Activity) error {
	// TODO send verify code
	// TODO check send limit
	exists, err := r.repo.UserRepo.ExistsByPhone(ctx, phone)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not check user exists in  db. err : %s", err.Error()))
	} else if exists {
		return errors.New("phone taken before")
	}

	err = r.repo.ConfirmCodeRepo.Create(phone)
	if err := r.repo.ActivityRepo.Create(ac); err != nil {
		panic(err)
	}
	return err
}

func (r *register) Verify(ctx context.Context, phone string, code string, ac *model.Activity) (*Auth, error) {
	// TODO check limits
	exists, err := r.repo.UserRepo.ExistsByPhone(ctx, phone)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not check user exists in  db. err : %s", err.Error()))
	}

	if exists {
		return nil, errors.New("user verified!")
	}

	if err := r.checkConfirmCode(phone, code); err != nil {
		return nil, err
	}

	user := r.createUser(ctx, phone)
	refreshToken, token := r.generateToken(ctx, user)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if err := r.repo.ConfirmCodeRepo.Delete(phone); err != nil {
			panic(fmt.Sprintf("internal error, can not delete confirm code. err : %s", err.Error()))
		}
		wg.Done()
	}()

	go func() {
		if err := r.repo.ActivityRepo.Create(ac); err != nil {
			panic(fmt.Sprintf("internal error, can not create activity log. err : %s", err.Error()))
		}
		wg.Done()
	}()
	wg.Wait()

	return &Auth{
		Token: token,
		RefreshToken: model.RefreshToken{
			Token: refreshToken.Token,
		},
		ExpiresIn: 3600, // 1 hour
		ID:        user.UUID,
	}, nil
}

func (r *register) createUser(ctx context.Context, phone string) *model.User {
	lastSeen := time.Now()
	user := &model.User{
		Phone:        &phone,
		LastSignInAt: &lastSeen,
		IsActive:     true,
	}

	if _, err := r.repo.UserRepo.Create(ctx, user); err != nil {
		panic(fmt.Sprintf("internal error, can not find token. err : %s", err.Error()))
	}

	return user
}

func (r *register) generateToken(ctx context.Context, user *model.User) (*model.RefreshToken, string) {
	refreshToken, err := r.repo.TokenRepo.Create(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	token, err := encrypt.GenerateAccessToken(user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	return refreshToken, token
}

func (r *register) checkConfirmCode(phone string, code string) error {
	confirm, err := r.repo.ConfirmCodeRepo.Find(phone)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not find confirm code. err : %s", err.Error()))
	}

	if confirm == nil || !encrypt.Check(confirm.Hash, code) {
		return errors.New("confirm code is invalid")
	}

	return nil
}
