package domain

import "context"

type RegisterService interface {
	Request(ctx context.Context) (string, error)
	Verify(ctx context.Context) (string, error)
}
