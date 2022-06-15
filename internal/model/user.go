package model

import (
	"database/sql"
	"gorm.io/gorm"

	"github.com/bxcodec/faker/v3"
)

type User struct {
	ID           string         `gorm:"primaryKey" gorm:"default:uuid_generate_v3()" faker:"uuid_digit"`
	Phone        sql.NullString `faker:"phone_number"`
	Email        sql.NullString `faker:"email"`
	Username     sql.NullString `faker:"username"`
	Password     sql.NullString `faker:"-" `
	OTPKey       sql.NullString `faker:"-" `
	OTPValue     sql.NullString `faker:"-" `
	ConfirmedAt  sql.NullTime   `faker:"-"`
	IsActive     bool           `faker:"-"`
	LastSignInAt sql.NullTime   `faker:"-"`
	CreatedAt    sql.NullTime   `gorm:"autoCreateTime" faker:"-"`
	UpdatedAt    sql.NullTime   `gorm:"autoUpdateTime" faker:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" faker:"-"`
}

type UserResource struct {
	Phone string `json:"phone"`
	UUID  string `json:"uuid"`
}

func (u User) SeedUser() (*User, error) {
	user := User{
		IsActive: true,
	}
	err := faker.FakeData(&user)
	return &user, err
}

//
//func (u *User) BeforeCreate(tx *gorm.DB) error {
//	u.UUID = uuid.New().String()
//	return nil
//}
