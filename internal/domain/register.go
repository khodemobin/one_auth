package domain

import "context"

type RegisterService interface {
	RegisterRequest(ctx context.Context, phone string) error
	RegisterVerify(ctx context.Context, phone string, code int) (*Login, error)
}

type RegisterRequest struct {
	Phone string `json:"phone" validate:"required,min=11,max=11,number"`
}

type RegisterVerifyRequest struct {
	Phone string `json:"phone" validate:"required,min=11,max=11,number"`
	Code  int    `json:"code" validate:"required,min=6,max=7,number"`
}
