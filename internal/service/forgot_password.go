package service

import (
	"context"
	"github.com/khodemobin/pilo/auth/internal/http/request"
	"github.com/khodemobin/pilo/auth/internal/repository"
)

type forgotPassword struct {
}

func NewForgotPasswordService(repo *repository.Repository) ForgotPasswordService {
	return &forgotPassword{}
}

func (f forgotPassword) Request(ctx context.Context, req request.ForgotPasswordRequest) {
	//TODO implement me
	panic("implement me")
}

func (f forgotPassword) Confirm(ctx context.Context, req request.ForgotPasswordConfirm) {
	//TODO implement me
	panic("implement me")
}
