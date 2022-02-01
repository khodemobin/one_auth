package domain

import "context"

type SampleService interface {
	Sample(ctx context.Context) (string, error)
}

type SampleRepository interface {
	Sample(ctx context.Context) (*int64, error)
}
