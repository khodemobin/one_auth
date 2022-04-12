package repository

import (
	"github.com/khodemobin/pilo/auth/internal/domain"
)

type Repository struct {
	UserRepo        domain.UserRepository
	TokenRepo       domain.TokenRepository
	ConfirmCodeRepo domain.ConfirmCodeRepository
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
