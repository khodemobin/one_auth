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
	ID           uint           `gorm:"primarykey" faker:"-" json:"-"`
	UUID         string         `db:"uuid" faker:"uuid_digit"`
	Phone        string         `json:"phone" db:"phone" faker:"phone_number"`
	Password     *string        `json:"-" db:"password" faker:"password" `
	ConfirmedAt  *time.Time     `json:"-" db:"confirmed_at" faker:"-"`
	Role         *string        `json:"-" db:"role" faker:"-"`
	Status       int            `json:"-" db:"status" faker:"-"`
	IsSuperAdmin bool           `json:"-" db:"is_super_admin" faker:"-"`
	LastSignInAt *time.Time     `json:"-" db:"last_sign_in_at" faker:"-"`
	CreatedAt    time.Time      `json:"-" faker:"-"`
	UpdatedAt    time.Time      `json:"-" faker:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index" faker:"-"`
	Tokens       []RefreshToken `json:"-" faker:"-"`
}

func (u User) SeedUser() (*User, error) {
	user := User{
		Status: 1,
	}
	err := faker.FakeData(&user)
	return &user, err
}
