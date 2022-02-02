package domain

import (
	"context"
	"fmt"
)

type Auth struct {
	Token string
}

type AuthError struct {
	Err error
}

func (r *AuthError) Error() string {
	return fmt.Sprintf("invalid credentials", r.Err)
}

type AuthService interface {
	Login(ctx context.Context, phone, password string) (*Auth, error)
	Check(ctx context.Context) (*Auth, error)
}

type AuthRepository interface {
	HasAccess(*Auth) error
}
