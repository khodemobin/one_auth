package repository

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo        UserRepository
	TokenRepo       TokenRepository
	ConfirmCodeRepo ConfirmCodeRepository
	ActivityRepos   ActivityRepository
}

func NewRepository() *Repository {
	up := NewUserRepo()
	tp := NewTokenRepo()
	cp := NewConfirmCodeRepo()
	ap := NewActivityRepo()
	return &Repository{
		UserRepo:        up,
		TokenRepo:       tp,
		ConfirmCodeRepo: cp,
		ActivityRepos:   ap,
	}
}

type ActivityRepository interface {
	Create(ac *model.Activity) error
}

type ConfirmCodeRepository interface {
	Create(phone string) error
	Find(phone string) (*model.ConfirmCode, error)
	Delete(phone string) error
}

type TokenRepository interface {
	Create(ctx context.Context, user *model.User) (*model.RefreshToken, error)
	Find(ctx context.Context, token string) (*model.RefreshToken, error)
	Revoke(ctx context.Context, token *model.RefreshToken) error
}

type UserRepository interface {
	FindByUUID(ctx context.Context, uuid string, status int) (*model.User, error)
	FindByID(ctx context.Context, id uint, status int) (*model.User, error)
	FindByPhone(ctx context.Context, phone string, status int) (*model.User, error)
	UpdateLastSeen(ctx context.Context, user *model.User) error
	CreateOrUpdate(ctx context.Context, user *model.User) error
}

func checkError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, redis.Nil) {
		return app.ErrNotFound
	}

	return err
}
