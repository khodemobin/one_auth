package domain

import (
	"context"
)

type Login struct {
	ID           string `json:"id"`
	Token        string `json:"token"`
	RefreshToken string `json:",omitempty"`
	ExpiresIn    int    `json:"expiresIn"`
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,min=11,max=11,number"`
	Password string `json:"password" validate:"required,min=5"`
}

type LoginService interface {
	Login(ctx context.Context, phone, password string, meta interface{}) (*Login, error)
	Logout(ctx context.Context, token string) error
}
