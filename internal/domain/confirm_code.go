package domain

import "time"

type ConfirmCode struct {
	Phone     string
	Hash      string
	ExpiresIn time.Duration
}

type ConfirmCodeRepository interface {
	CreateConfirmCode(phone string) error
	FindConfirmCode(phone string) (*ConfirmCode, error)
	DeleteConfirmCode(phone string) error
}
