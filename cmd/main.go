package main

import (
	"context"
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/BalamutDiana/grps_server/internal/config"
	"github.com/BalamutDiana/grps_server/internal/repository"
	"github.com/BalamutDiana/grps_server/internal/service"
	products "github.com/BalamutDiana/grps_server/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getRowsFromCSV(url string, service *service.Product) {

	data, err := readCSVFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}

	for idx, row := range data {
		if idx == 0 {
			continue
		}

		if idx == 0 {
			break
		}

		key := row[0]
		value, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(key, value)
		var item products.Product

		item.Name = key
		item.Price = value
		item.ChangesCount = 1
		item.Timestamp = time.Now()

		service.Insert(context.TODO(), item)
	}
}

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

	//url := "http://164.92.251.245:8080/api/v1/products/"
	//getRowsFromCSV(url, productsService)

	var item products.Product
	item, err = productsService.GetByName(context.TODO(), "Unbranded Plastic Chicken")
	item.Price = 453453
	item.Timestamp = time.Now()

	err = productsService.UpdateByName(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}
}
