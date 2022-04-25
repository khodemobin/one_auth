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
	ActivityRepo    ActivityRepository
	AccessTokenRepo AccessTokenRepository
}

func NewRepository() *Repository {
	up := NewUserRepo()
	tp := NewTokenRepo()
	cp := NewConfirmCodeRepo()
	ap := NewActivityRepo()
	at := NewAccessTokenRepo()
	return &Repository{
		UserRepo:        up,
		TokenRepo:       tp,
		ConfirmCodeRepo: cp,
		ActivityRepo:    ap,
		AccessTokenRepo: at,
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

type AccessTokenRepository interface {
	AddToBlacklist(token string) error
	ExistsInBlackList(token string) (bool, error)
}

type UserRepository interface {
	Find(ctx context.Context, column string, value string) (*model.User, error)
	FindActive(ctx context.Context, column string, value string) (*model.User, error)
	UpdateLastSeen(ctx context.Context, user *model.User) error
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id uint) (bool, error)
	ExistsByUUID(ctx context.Context, uuid string) (bool, error)
}

func checkError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, redis.Nil) {
		return app.ErrNotFound
	}

	return err
}
