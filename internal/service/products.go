package service

import (
	"context"
	"strconv"
	"time"

	products "github.com/BalamutDiana/grps_server/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Repository interface {
	Insert(ctx context.Context, item products.Product) error
	GetByName(ctx context.Context, name string) (products.Product, error)
	UpdateByName(ctx context.Context, prod products.Product) error
	List(ctx context.Context, paging products.PagingParams, sorting products.SortingParams) ([]products.Product, error)
}

type Product struct {
	repo Repository
}

func NewProduct(repo Repository) *Product {
	return &Product{
		repo: repo,
	}
}

func (s *Product) List(ctx context.Context, req *products.ListRequest) (*products.ListResponse, error) {
	paging := products.PagingParams{
		Offset: int(req.GetPagingOffset()),
		Limit:  int(req.GetPagingLimit()),
	}
	sorting := products.SortingParams{
		Field: req.SortField,
		Asc:   req.SortAsc,
	}

	items, err := s.repo.List(ctx, paging, sorting)
	if err != nil {
		return nil, err
	}
	var sorted_products []*products.ProductItem

	for _, x := range items {
		var sorted_product products.ProductItem
		sorted_product.Name = x.Name
		sorted_product.Price = int32(x.Price)
		sorted_product.Count = int32(x.ChangesCount)
		sorted_product.Timestamp = timestamppb.New(x.Timestamp)
		sorted_products = append(sorted_products, &sorted_product)
	}

	return &products.ListResponse{
		Product: sorted_products,
	}, nil

}

func (s *Product) Fetch(ctx context.Context, req *products.FetchRequest) (*products.FetchResponse, error) {

	url := req.Url
	data, err := readCSVFromUrl(url)
	if err != nil {
		return nil, err
	}

	for idx, row := range data {
		if idx == 0 {
			continue
		}

		name := row[0]
		price, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, err
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
					return nil, err
				}
				continue
			} else {
				return nil, err
			}
		}

		if prod.Price != item.Price {
			prod.Price = price
			prod.Timestamp = time.Now()

			if err = s.repo.UpdateByName(ctx, prod); err != nil {
				return nil, err
			}
		}
	}
	return &products.FetchResponse{
		Status: "OK",
	}, nil
}
