package service

import (
	"context"
	"log"
	"strconv"
	"time"

	products "github.com/BalamutDiana/grps_server/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Insert(ctx context.Context, item products.Product) error
	GetByName(ctx context.Context, name string) (products.Product, error)
	UpdateByName(ctx context.Context, prod products.Product) error
	List(ctx context.Context, paging products.PagingParams, sorting []products.SortingParams) ([]products.Product, error)
}

type Product struct {
	repo Repository
}

func NewProduct(repo Repository) *Product {
	return &Product{
		repo: repo,
	}
}

func (s *Product) List(ctx context.Context, paging products.PagingParams, sorting []products.SortingParams) ([]products.Product, error) {
	return s.repo.List(ctx, paging, sorting)
}

func (s *Product) Fetch(ctx context.Context, url string) error {

	data, err := readCSVFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}

	for idx, row := range data {
		if idx == 0 {
			continue
		}

		// if idx == 30 {
		// 	break
		// }

		name := row[0]
		price, err := strconv.Atoi(row[1])
		if err != nil {
			return err
		}

		var prod products.Product

		item, err := s.repo.GetByName(ctx, name)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				prod.Name = name
				prod.Price = price
				prod.ChangesCount = 1
				prod.Timestamp = time.Now()

				if err = s.repo.Insert(ctx, prod); err != nil {
					return err
				}
				continue
			} else {
				return err
			}
		}

		if prod.Price != item.Price {
			prod.Price = price
			prod.Timestamp = time.Now()

			if err = s.repo.UpdateByName(ctx, prod); err != nil {
				return err
			}
		}
	}
	return nil
}
