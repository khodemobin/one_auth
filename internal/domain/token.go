package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	ID        uint           `gorm:"primarykey" faker:"-"`
	Token     string         `db:"token"`
	UserID    uint           `db:"user_id"`
	Revoked   bool           `db:"revoked"`
	CreatedAt time.Time      `json:"created_at" faker:"-"`
	UpdatedAt time.Time      `json:"updated_at" faker:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" faker:"-"`
}

type TokenRepository interface {
	Create(ctx context.Context, ttl int, user *User) (*Token, error)
	Revoke(ctx context.Context, token *Token) error
}
