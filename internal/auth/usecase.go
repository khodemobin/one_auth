//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package auth

import (
	"context"
	"github.com/khodemobin/pilo/auth/internal/http/request"
	"github.com/khodemobin/pilo/auth/internal/model"
)

type Auth struct {
	ID           string             `json:"id"`
	Token        string             `json:"token"`
	RefreshToken model.RefreshToken `json:"-"`
	ExpiresIn    int                `json:"expiresIn"`
}

type UseCase interface {
	Login(ctx context.Context, req request.LoginRequest) (*Auth, error)
	Logout(ctx context.Context, accessToken, token string) error
	Refresh(ctx context.Context, tokenString string) (*Auth, error)
}
