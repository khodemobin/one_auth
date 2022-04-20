package service

import (
	"context"

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
	LoginService    LoginService
	RegisterService RegisterService
	UserService     UserService
	RefreshService  RefreshTokenService
}

func NewService(repo *repository.Repository) *Service {
	login := NewLoginService(repo)
	register := NewRegisterService(repo)
	user := NewUserService(repo)
	refresh := NewRefreshTokenService(repo)

	return &Service{
		LoginService:    login,
		RegisterService: register,
		UserService:     user,
		RefreshService:  refresh,
	}
}

type LoginService interface {
	Login(ctx context.Context, phone, password string, ac *model.Activity) (*Auth, error)
	Logout(ctx context.Context, token string, ac *model.Activity) error
}

type RegisterService interface {
	RegisterRequest(ctx context.Context, phone string, ac *model.Activity) error
	RegisterVerify(ctx context.Context, phone string, code string, ac *model.Activity) (*Auth, error)
}

type UserService interface {
	GetUser(ctx context.Context, uuid string, ac *model.Activity) (*model.User, error)
	UpdateUser(ctx context.Context, uuid string, user *model.User, ac *model.Activity) error
}

type RefreshTokenService interface {
	RefreshToken(ctx context.Context, tokenString string, ac *model.Activity) (*Auth, error)
}