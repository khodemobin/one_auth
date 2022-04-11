package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID        uint           `gorm:"primarykey" faker:"-"`
	Token     string         `db:"token"`
	UserID    uint           `db:"user_id"`
	CreatedAt time.Time      `json:"created_at" faker:"-"`
	UpdatedAt time.Time      `json:"updated_at" faker:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" faker:"-"`
}

type TokenRepository interface {
	CreateToken(ctx context.Context, ttl int, user *User) (*Token, error)
	RevokeToken(ctx context.Context, token *Token) error
}

func (Token) TableName() string {
	return "refresh_tokens"
}
