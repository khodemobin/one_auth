package model

import (
	"database/sql"
)

type RefreshToken struct {
	ID        uint `faker:"-"`
	Token     string
	UserID    string
	Revoked   bool
	CreatedAt sql.NullTime `gorm:"autoCreateTime" faker:"-"`
	UpdatedAt sql.NullTime `gorm:"autoUpdateTime" faker:"-"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
