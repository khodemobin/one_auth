package domain

import "context"

type RecoveryService interface {
	Request(ctx context.Context) (string, error)
	Verify(ctx context.Context) (string, error)
}
