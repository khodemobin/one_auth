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
	ID           uint           `gorm:"primarykey" db:"id" faker:"-"`
	UUID         string         `db:"uuid" faker:"uuid_digit"`
	Phone        string         `db:"phone" faker:"phone_number"`
	Password     *string        `db:"password" faker:"password" `
	ConfirmedAt  *time.Time     `db:"confirmed_at" faker:"-"`
	Role         *string        `db:"role" faker:"-"`
	Status       int            `db:"status" faker:"-"`
	IsSuperAdmin bool           `db:"is_super_admin" faker:"-"`
	LastSignInAt *time.Time     `db:"last_sign_in_at" faker:"-"`
	CreatedAt    time.Time      `faker:"-"`
	UpdatedAt    time.Time      `faker:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" faker:"-"`
	Tokens       []RefreshToken `faker:"-"`
}

type UserResource struct {
	Phone string `json:"phone"`
	UUID  string `json:"uuid"`
}

func (u User) SeedUser() (*User, error) {
	user := User{
		Status: 1,
	}
	err := faker.FakeData(&user)
	return &user, err
}
