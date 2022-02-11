package repository

import (
	"context"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"gorm.io/gorm"
	"time"
)

type token struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewTokenRepo(db *gorm.DB, cache cache.Cache) domain.TokenRepository {
	return &token{
		db,
		cache,
	}
}

func (t token) CreateToken(ctx context.Context, ttl int, user *domain.User) (*domain.Token, error) {
	token, err := encrypt.GenerateAccessToken(user, time.Second*time.Duration(ttl))
	if err != nil {
		return nil, err
	}

	tokenModel := &domain.Token{
		Token:   token,
		UserID:  user.ID,
		Revoked: false,
	}

	err = t.db.Create(&tokenModel).Error
	return tokenModel, err
}

func (t token) RevokeToken(ctx context.Context, token *domain.Token) error {
	return nil
}
