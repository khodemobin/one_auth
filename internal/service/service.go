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
	AuthService           AuthService
	ForgotPasswordService ForgotPasswordService
	VerifyService         VerifyService
	UserService           UserService

	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo),
		//ForgotPasswordService: NewForgotPasswordService(repo),
		//VerifyService:         NewVerifyService(repo),
		//UserService:           NewUserService(repo),
		Repo: repo,
	}
}

type AuthService interface {
	Login(ctx context.Context, req request.LoginRequest) (*Auth, error)
	Logout(ctx context.Context, accessToken, token string) error
	Refresh(ctx context.Context, tokenString string) (*Auth, error)
}

type ForgotPasswordService interface {
	Request(ctx context.Context, req request.ForgotPasswordRequest)
	Confirm(ctx context.Context, req request.ForgotPasswordConfirm)
}

type VerifyService interface {
	Request(ctx context.Context, req request.VerifyRequest) error
	Verify(ctx context.Context, req request.VerifyConfirmRequest) (*Auth, error)
}

type UserService interface {
	Me(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, id string, req request.UserUpdateRequest) error
	Create(ctx context.Context, req request.UserCreateRequest) error
	Delete(ctx context.Context, id string) error
}
