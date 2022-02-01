package domain

import (
	"time"
)

type User struct {
	ID        int        `json:"id" db:"id"`
	Phone     string     `json:"phone" db:"phone"`
	Password  *string    `json:"password" db:"password"`
	Status    string     `json:"status" db:"status"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type UserRepository interface{}
