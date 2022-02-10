package domain

import (
	"context"
	"gorm.io/gorm"
	"time"

	"github.com/bxcodec/faker/v3"
)

const (
	USER_STATUS_ACTIVE    = 1
	USER_STATUS_IN_ACTIVE = 0
	USER_STATUS_PENDDING  = 2
)

type User struct {
	ID                 uint           `gorm:"primarykey" faker:"-"`
	Phone              string         `json:"phone" db:"phone" faker:"phone_number"`
	Password           *string        `json:"password" db:"password" faker:"password" `
	ConfirmationToken  *string        `json:"confirmation_token" db:"confirmation_token" faker:"jwt"`
	ConfirmationSentAt *time.Time     `json:"confirmation_sent_at" db:"confirmation_sent_at" faker:"-"`
	ConfirmedAt        *time.Time     `json:"confirmed_at" db:"confirmed_at" faker:"-"`
	RecoveryToken      *string        `json:"recovery_token" db:"recovery_token" faker:"jwt"`
	RecoverySentAt     *time.Time     `json:"recovery_sent_at" db:"recovery_sent_at" faker:"-"`
	PhoneChangeToken   *string        `json:"phone_change_token" db:"phone_change_token" faker:"jwt"`
	PhoneChange        *string        `json:"phone_change" db:"phone_change" faker:"phone_number"`
	PhoneChangeSentAt  *time.Time     `json:"phone_change_sent_at" db:"phone_change_sent_at" faker:"-"`
	Role               *string        `json:"role" db:"role" faker:"-"`
	Status             int            `json:"status" db:"status" faker:"-"`
	IsSuperAdmin       bool           `json:"is_super_admin" db:"is_super_admin" faker:"-"`
	LastSignInAt       *time.Time     `json:"last_sign_in_at" db:"last_sign_in_at" faker:"-"`
	CreatedAt          time.Time      `json:"created_at" faker:"-"`
	UpdatedAt          time.Time      `json:"updated_at" faker:"-"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"index" faker:"-"`
	Tokens             []Token        `faker:"-"`
}

type UserRepository interface {
	FindUserById(ctx context.Context, id int) (*User, error)
	FindUserByPhone(ctx context.Context, phone string) (*User, error)
	UpdateUserLastSeen(ctx context.Context, user *User) error
}

func (u User) SeedUser() (*User, error) {
	user := User{
		Status: 1,
	}
	err := faker.FakeData(&user)
	return &user, err
}
