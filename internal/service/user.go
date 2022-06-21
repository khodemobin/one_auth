package service

import (
	"context"
	"fmt"
	"github.com/khodemobin/pilo/auth/internal/http/request"
	"github.com/khodemobin/pilo/auth/pkg/utils/encrypt"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
)

type user struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) UserService {
	return &user{
		repo: repo,
	}
}

func (u *user) Me(ctx context.Context, id string) (*model.User, error) {
	user, err := u.repo.UserRepo.FindActive(ctx, "id", id)
	if errors.Is(err, app.ErrNotFound) {
		return nil, errors.New("user not found")
	} else if err != nil {
		panic(fmt.Sprintf("internal error, can not find user from db. err : %s", err.Error()))
	}

	return user, err
}

func (u *user) Update(ctx context.Context, id string, req request.UserUpdateRequest) error {
	if req.Password != req.ConfirmPassword {
		return errors.New("password and confirm password are not equals")
	}

	user, err := u.repo.UserRepo.FindActive(ctx, "id", id)
	if errors.Is(err, app.ErrNotFound) {
		return errors.New("user not found")
	} else if err != nil {
		panic(fmt.Sprintf("internal error, can not find user from db. err : %s", err.Error()))
	}

	p, err := encrypt.Hash(req.Password)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not encrypt password. err : %s", err.Error()))
	}

	user.Password.String = p
	_, err = u.repo.UserRepo.Update(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not update user password from db. err : %s", err.Error()))
	}

	return nil
}

func (u *user) Create(ctx context.Context, req request.UserCreateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u *user) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
