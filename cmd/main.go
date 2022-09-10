package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BalamutDiana/grps_server/internal/config"
	"github.com/BalamutDiana/grps_server/internal/repository"
	"github.com/BalamutDiana/grps_server/internal/service"
	products "github.com/BalamutDiana/grps_server/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	productsRepo := repository.NewProducts(db)
	productsService := service.NewProduct(productsRepo)

	url := "http://164.92.251.245:8080/api/v1/products/"
	if err = productsService.Fetch(ctx, url); err != nil {
		log.Fatal(err)
	}

	paging := products.PagingParams{
		Offset: 0,
		Limit:  10,
	}

	sorting := products.NewProductsSortingParams()

	sorting1 := products.SortingParams{
		Field: sorting.Price,
		Asc:   false,
	}

	var sortingList []products.SortingParams
	sortingList = append(sortingList, sorting1)

	var item []products.Product
	item, err = productsService.List(ctx, paging, sortingList)

	for _, x := range item {
		fmt.Println(x)
	}
}
