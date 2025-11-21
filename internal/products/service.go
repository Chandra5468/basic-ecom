package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
	// repository
}

func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}

func NewService() Service {
	return &svc{}
}
