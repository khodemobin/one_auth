package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/domain"
)

type Repository struct {
	Sample domain.SampleRepository
}

func NewRepository(db *sqlx.DB, cache cache.Cache) *Repository {
	s := NewSampleRepo(db, cache)
	return &Repository{
		Sample: s,
	}
}
