package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/domain"
)

type Repository struct {
	UserRepo domain.UserRepository
}

func NewRepository(db *sqlx.DB, cache cache.Cache) *Repository {
	up := NewUserRepo(db, cache)
	return &Repository{
		UserRepo: up,
	}
}
