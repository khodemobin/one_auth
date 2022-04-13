package repository

import (
	"context"

	"github.com/khodemobin/pilo/auth/internal/model"
)

type Repository struct {
	UserRepo        UserRepository
	TokenRepo       TokenRepository
	ConfirmCodeRepo ConfirmCodeRepository
}

func NewRepository() *Repository {
	up := NewUserRepo()
	tp := NewTokenRepo()
	cp := NewConfirmCodeRepo()
	return &Repository{
		UserRepo:        up,
		TokenRepo:       tp,
		ConfirmCodeRepo: cp,
	}
}

type ActivityRepository interface {
	CreateActivity(ac model.Activity) error
}

type ConfirmCodeRepository interface {
	CreateConfirmCode(phone string) error
	FindConfirmCode(phone string) (*model.ConfirmCode, error)
	DeleteConfirmCode(phone string) error
}

type TokenRepository interface {
	CreateToken(ctx context.Context, ttl int, user *model.User) (*model.RefreshToken, error)
	RevokeToken(ctx context.Context, token *model.RefreshToken) error
}

type UserRepository interface {
	FindUserById(ctx context.Context, id int, status int) (*model.User, error)
	FindUserByPhone(ctx context.Context, phone string, status int) (*model.User, error)
	UpdateUserLastSeen(ctx context.Context, user *model.User) error
	CreateOrUpdateUser(ctx context.Context, user *model.User) error
}
