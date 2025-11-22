package products

import (
	"context"

	repo "github.com/Chandra5468/basic-ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	// repository
	// repo repo.Queries this is struct. But always prefer interfaces
	repo repo.Querier // this is interface
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	// products, err := s.repo.ListProducts(ctx)
	return s.repo.ListProducts(ctx)
}
