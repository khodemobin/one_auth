package domain

import "time"

type ConfirmCode struct {
	Phone     string
	Hash      string
	ExpiresIn time.Duration
}

type ConfirmCodeRepository interface {
	Store(phone string, confirmCode *ConfirmCode) error
	Find(phone string) (*ConfirmCode, error)
}
