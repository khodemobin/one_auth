package repository

import (
	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/domain"
)

type authRepo struct {
	cache *cache.Cache
}

func NewAuthRepository(cache *cache.Cache) domain.AuthRepository {
	return &authRepo{
		cache: cache,
	}
}

func (a *authRepo) HasAccess(*domain.Auth) error {
	return nil
}
