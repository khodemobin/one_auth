package domain

import "context"

type AuthService interface {
	Login(ctx context.Context) (string, error)
	Check(ctx context.Context) (string, error)
}
