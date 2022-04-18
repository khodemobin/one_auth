package repository

import (
	"context"

	"github.com/khodemobin/pilo/auth/internal/model"
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
	CreateActivity(ac *model.Activity) error
}

type ConfirmCodeRepository interface {
	CreateConfirmCode(phone string) error
	FindConfirmCode(phone string) (*model.ConfirmCode, error)
	DeleteConfirmCode(phone string) error
}

type TokenRepository interface {
	CreateToken(ctx context.Context, user *model.User) (*model.RefreshToken, error)
	FindToken(ctx context.Context, token string) (*model.RefreshToken, error)
	RevokeToken(ctx context.Context, token *model.RefreshToken) error
}

type UserRepository interface {
	FindUserByUUID(ctx context.Context, uuid string, status int) (*model.User, error)
	FindUserByPhone(ctx context.Context, phone string, status int) (*model.User, error)
	UpdateUserLastSeen(ctx context.Context, user *model.User) error
	CreateOrUpdateUser(ctx context.Context, user *model.User) error
}
