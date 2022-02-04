package domain

import (
	"context"
)

type Auth struct {
	Token string
}

type AuthService interface {
	Login(ctx context.Context, phone, password string) (*Auth, error)
	Check(ctx context.Context) (*Auth, error)
}

type AuthRepository interface {
	HasAccess(*Auth) error
}
