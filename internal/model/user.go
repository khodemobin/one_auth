package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/bxcodec/faker/v3"
)

const (
	USER_STATUS_ACTIVE    = 1
	USER_STATUS_IN_ACTIVE = 0
)

type User struct {
	ID           uint           `gorm:"primarykey" faker:"-"`
	UUID         string         `gorm:"uuid" faker:"uuid_digit"`
	Phone        string         `json:"phone" db:"phone" faker:"phone_number"`
	Password     *string        `json:"password" db:"password" faker:"password" `
	ConfirmedAt  *time.Time     `json:"confirmed_at" db:"confirmed_at" faker:"-"`
	Role         *string        `json:"role" db:"role" faker:"-"`
	Status       int            `json:"status" db:"status" faker:"-"`
	IsSuperAdmin bool           `json:"is_super_admin" db:"is_super_admin" faker:"-"`
	LastSignInAt *time.Time     `json:"last_sign_in_at" db:"last_sign_in_at" faker:"-"`
	CreatedAt    time.Time      `json:"created_at" faker:"-"`
	UpdatedAt    time.Time      `json:"updated_at" faker:"-"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index" faker:"-"`
	Tokens       []RefreshToken `faker:"-"`
}

func (u User) SeedUser() (*User, error) {
	user := User{
		Status: 1,
	}
	err := faker.FakeData(&user)
	return &user, err
}
