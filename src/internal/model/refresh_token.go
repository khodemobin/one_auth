package model

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uint           `gorm:"primarykey" db:"id" faker:"-"`
	Token     string         `db:"token"`
	UserID    uint           `db:"user_id"`
	Revoked   bool           `db:"revoked"`
	CreatedAt time.Time      `json:"created_at" faker:"-"`
	UpdatedAt time.Time      `json:"updated_at" faker:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" faker:"-"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
