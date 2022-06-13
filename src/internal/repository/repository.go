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
	UserRepo      UserRepository
	TokenRepo     TokenRepository
	ActivityRepo  ActivityRepository
	BlackListRepo BlackListRepository
}

func NewRepository() *Repository {
	up := NewUserRepo()
	tp := NewTokenRepo()
	ap := NewActivityRepo()
	blt := NewBlackListRepo()
	return &Repository{
		UserRepo:      up,
		TokenRepo:     tp,
		ActivityRepo:  ap,
		BlackListRepo: blt,
	}
}

type ActivityRepository interface {
	Create(ac *model.Activity) error
}

type TokenRepository interface {
	Create(ctx context.Context, user *model.User) (*model.RefreshToken, error)
	Find(ctx context.Context, token string) (*model.RefreshToken, error)
	Revoke(ctx context.Context, token string) error
}

type BlackListRepository interface {
	Create(token string) error
	Exists(token string) (bool, error)
}

type UserRepository interface {
	Find(ctx context.Context, column string, value string) (*model.User, error)
	FindActive(ctx context.Context, column string, value string) (*model.User, error)
	UpdateLastSeen(ctx context.Context, user *model.User) error
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id uint) (bool, error)
}

func checkError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, redis.Nil) {
		return app.ErrNotFound
	}

	return err
}
