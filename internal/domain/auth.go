package domain

import (
	"context"
)

type Auth struct {
	UserID       uint   `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:",omitempty"`
	ExpiresIn    int    `json:"expiresIn"`
}

type AuthRequest struct {
	Phone    string `json:"phone" validate:"required,min=11,max=11,number"`
	Password string `json:"password" validate:"required,min=5"`
}

type AuthService interface {
	Login(ctx context.Context, phone, password string) (*Auth, error)
	Check(ctx context.Context) (*Auth, error)
}
