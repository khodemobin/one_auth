package model

import "time"

type ConfirmCode struct {
	Phone     string
	Hash      string
	ExpiresIn time.Duration
}
