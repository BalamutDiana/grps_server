package repository

import (
	"context"

	product "github.com/BalamutDiana/grps_server/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Products struct {
	db *mongo.Database
}

func NewProducts(db *mongo.Database) *Products {
	return &Products{
		db: db,
	}
}

func (r *Products) Insert(ctx context.Context, item product.Product) error {
	_, err := r.db.Collection("products").InsertOne(ctx, item)

	return err
}

func (r *Products) GetByName(ctx context.Context, name string) (product.Product, error) {
	var prod product.Product
	filter := bson.D{{Key: "name", Value: name}}

	err := r.db.Collection("products").FindOne(ctx, filter).Decode(&prod)
	if err != nil {
		return prod, err
	}
	return prod, nil
}
