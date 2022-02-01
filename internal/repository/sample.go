package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/domain"
)

type sampleRepo struct {
	db    *sqlx.DB
	cache cache.Cache
}

func NewSampleRepo(db *sqlx.DB, cache cache.Cache) domain.SampleRepository {
	return &sampleRepo{
		db:    db,
		cache: cache,
	}
}

func (r *sampleRepo) Sample(ctx context.Context) (*int64, error) {
	return nil, nil
}
