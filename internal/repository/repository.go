package repository

import (
	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo  domain.UserRepository
	TokenRepo domain.TokenRepository
}

func NewRepository(db *gorm.DB, cache cache.Cache, cfg *config.Config) *Repository {
	up := NewUserRepo(db, cache)
	tp := NewTokenRepo(db, cache, cfg)
	return &Repository{
		UserRepo:  up,
		TokenRepo: tp,
	}
}
