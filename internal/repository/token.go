package repository

import (
	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"gorm.io/gorm"
	"time"
)

type token struct {
	db    *gorm.DB
	cache cache.Cache
	cfg   *config.Config
}

func NewTokenRepo(db *gorm.DB, cache cache.Cache, cfg *config.Config) domain.TokenRepository {
	return &token{
		db,
		cache,
		cfg,
	}
}

func (t token) Create(ttl int, user *domain.User) (*domain.Token, error) {
	token, err := encrypt.GenerateAccessToken(user, time.Second*time.Duration(ttl), t.cfg.App.JwtSecret)
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

func (t token) Revoke(token *domain.Token) error {
	return nil
}
