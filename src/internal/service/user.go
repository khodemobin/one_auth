package service

import (
	"context"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type user struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) UserService {
	return &user{
		repo: repo,
	}
}

// GetUser implements UserService
func (u *user) Me(ctx context.Context, uuid string) (*model.User, error) {
	user, err := u.repo.UserRepo.FindByUUID(ctx, uuid, model.USER_STATUS_ACTIVE)
	if errors.Is(err, app.ErrNotFound) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can not find user from db. err : %s", err.Error()))
	}

	return user, err
}

// Update implements UserService
func (u *user) Update(ctx context.Context, uuid string, password string, confirm string, ac *model.Activity) error {
	if password != confirm {
		return errors.New("password and confirm password are not equals")
	}

	user, err := u.repo.UserRepo.FindByUUID(ctx, uuid, model.USER_STATUS_ACTIVE)
	if errors.Is(err, app.ErrNotFound) {
		return errors.New("user not found")
	}

	if err != nil {
		panic(fmt.Sprintf("internal error, can not find user from db. err : %s", err.Error()))
	}

	p, err := encrypt.Hash(password)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not encrypt password. err : %s", err.Error()))
	}

	user.Password = &p
	err = u.repo.UserRepo.CreateOrUpdate(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not update user password from db. err : %s", err.Error()))
	}

	return nil
}
