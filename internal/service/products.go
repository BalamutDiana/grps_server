package service

import (
	"context"

	products "github.com/BalamutDiana/grps_server/pkg/domain"
)

type Repository interface {
	Insert(ctx context.Context, item products.Product) error
	GetByName(ctx context.Context, name string) (products.Product, error)
	UpdateByName(ctx context.Context, prod products.Product) error
}

type Product struct {
	repo Repository
}

func NewProduct(repo Repository) *Product {
	return &Product{
		repo: repo,
	}
}

func (s *Product) Insert(ctx context.Context, req products.Product) error {
	return s.repo.Insert(ctx, req)
}

func (s *Product) GetByName(ctx context.Context, name string) (products.Product, error) {
	return s.repo.GetByName(ctx, name)
}

func (s *Product) UpdateByName(ctx context.Context, prod products.Product) error {
	return s.repo.UpdateByName(ctx, prod)
}
