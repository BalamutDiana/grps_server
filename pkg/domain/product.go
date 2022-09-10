package products

import (
	"time"
)

type Product struct {
	Name         string    `bson:"name"`
	Price        int       `bson:"price"`
	ChangesCount int       `bson:"changes_count"`
	Timestamp    time.Time `bson:"timestamp"`
}

type PagingParams struct {
	Offset int
	Limit  int
}

/*
fields along with their sorting order
1 is used for ascending order while -1 is used for descending order
*/

type ProductsSortingParams struct {
	Name         string
	Price        string
	ChangesCount string
	Timestamp    string
}

func NewProductsSortingParams() ProductsSortingParams {
	return ProductsSortingParams{
		Name:         "name",
		Price:        "price",
		ChangesCount: "changes_count",
		Timestamp:    "timestamp",
	}
}

type SortingParams struct {
	Field interface{}
	Asc   bool
}
