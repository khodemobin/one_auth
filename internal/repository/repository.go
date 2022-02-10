package repository

import (
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo  domain.UserRepository
	TokenRepo domain.TokenRepository
}

func NewRepository(db *gorm.DB, cache cache.Cache) *Repository {
	up := NewUserRepo(db, cache)
	tp := NewTokenRepo(db, cache)
	return &Repository{
		UserRepo:  up,
		TokenRepo: tp,
	}
}
