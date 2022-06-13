package service

import (
	"context"
	"github.com/khodemobin/pilo/auth/internal/http/request"

	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
)

type Auth struct {
	ID           string             `json:"id"`
	Token        string             `json:"token"`
	RefreshToken model.RefreshToken `json:"-"`
	ExpiresIn    int                `json:"expiresIn"`
}

type Service struct {
	AuthService     AuthService
	RegisterService RegisterService
	UserService     UserService
	RefreshService  RefreshTokenService

	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	auth := NewAuthService(repo)
	register := NewRegisterService(repo)
	user := NewUserService(repo)
	refresh := NewRefreshTokenService(repo)

	return &Service{
		AuthService:     auth,
		RegisterService: register,
		UserService:     user,
		RefreshService:  refresh,
		Repo:            repo,
	}
}

type AuthService interface {
	Login(ctx context.Context, request request.LoginRequest) (*Auth, error)
	Logout(ctx context.Context, accessToken, token string) error
	ForgotPassword()
}

//type RegisterService interface {

//}

type VerifyService interface {
	Request(ctx context.Context) error
	Verify(ctx context.Context, phone, code string) (*Auth, error)
}

type UserService interface {
	Me(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context) error
	Create(ctx context.Context) error
	Delete(ctx context.Context, uuid, password, confirm string) error
}

type RefreshTokenService interface {
	Refresh(ctx context.Context, tokenString string) (*Auth, error)
}
