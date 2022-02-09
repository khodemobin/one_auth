package domain

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token   string `db:"token"`
	UserID  uint   `db:"user_id"`
	Revoked bool   `db:"revoked"`
}

type TokenRepository interface {
	Create(ttl int, user *User) (*Token, error)
	Revoke(token *Token) error
}
