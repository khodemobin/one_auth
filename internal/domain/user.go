package domain

import (
	"time"

	"gorm.io/gorm"
)

const (
	USER_STATUS_ACTIVE    = 1
	USER_STATUS_IN_ACTIVE = 0
	USER_STATUS_PENDDING  = 2
)

type User struct {
	gorm.Model
	Phone              string     `json:"phone" db:"phone"`
	Password           *string    `json:"password" db:"password"`
	ConfirmationToken  *string    `json:"confirmation_token" db:"confirmation_token"`
	ConfirmationSentAt *time.Time `json:"confirmation_sent_at" db:"confirmation_sent_at"`
	ConfirmedAt        *time.Time `json:"confirmed_at" db:"confirmed_at"`
	RecoveryToken      *string    `json:"recovery_token" db:"recovery_token"`
	RecoverySentAt     *time.Time `json:"recovery_sent_at" db:"recovery_sent_at"`
	PhoneChangeToken   *string    `json:"phone_change_token" db:"phone_change_token"`
	PhoneChange        *string    `json:"phone_change" db:"phone_change"`
	PhoneChangeSentAt  *time.Time `json:"phone_change_sent_at" db:"phone_change_sent_at"`
	Role               *string    `json:"role" db:"role"`
	Status             int        `json:"status" db:"status"`
	IsSuperAdmin       bool       `json:"is_super_admin" db:"is_super_admin"`
	LastSignInAt       *time.Time `json:"last_sign_in_at" db:"last_sign_in_at"`
	Tokens             []Token
}

type UserRepository interface {
	FindUserById(id int) (*User, error)
	FindUserByPhone(phone string) (*User, error)
	UpdateUserLastSeen(user *User) (*User, error)
}
