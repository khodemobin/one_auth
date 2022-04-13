package service

import (
	"context"

	"github.com/khodemobin/pilo/auth/internal/repository"
)

type Auth struct {
	ID           string `json:"id"`
	Token        string `json:"token"`
	RefreshToken string `json:",omitempty"`
	ExpiresIn    int    `json:"expiresIn"`
}

type MetaData struct {
	Headers map[string]string
	IPs     []string
	Path    string
}

type Service struct {
	LoginService    LoginService
	RegisterService RegisterService
}

func NewService(repo *repository.Repository) *Service {
	login := NewLoginService(repo)
	register := NewRegisterService(repo)

	return &Service{
		LoginService:    login,
		RegisterService: register,
	}
}

type LoginService interface {
	Login(ctx context.Context, phone, password string, meta *MetaData) (*Auth, error)
	Logout(ctx context.Context, token string, meta *MetaData) error
}

type RegisterService interface {
	RegisterRequest(ctx context.Context, phone string, meta *MetaData) error
	RegisterVerify(ctx context.Context, phone string, code string, meta *MetaData) (*Auth, error)
}
