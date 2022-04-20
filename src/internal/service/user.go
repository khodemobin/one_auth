package service

import (
	"context"

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

// GetUser implements UserService
func (u *user) Me(ctx context.Context, uuid string, ac *model.Activity) (*model.User, error) {
	user, err := u.repo.UserRepo.FindByUUID(ctx, uuid, model.USER_STATUS_ACTIVE)
	if errors.Is(err, app.ErrNotFound) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		panic(err)
	}

	if err := u.repo.ActivityRepos.Create(ac); err != nil {
		panic(err)
	}

	return user, err
}

// UpdateUser implements UserService
func (u *user) Update(ctx context.Context, uuid string, user *model.User, ac *model.Activity) error {
	// err := u.repo.UserRepo.UpdatePassword()
	return nil
}
