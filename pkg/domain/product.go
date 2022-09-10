package products

import (
	"errors"
	"time"
)

const (
	SORTINGFIELD_NAME  = "name"
	SORTINGFIELD_PRICE = "price"
	SORTINGFIELD_COUNT = "changes_count"
	SORTINGFIELD_TIME  = "timestamp"
)

var (
	fields = map[string]ListRequest_SortingField{
		SORTINGFIELD_NAME:  ListRequest_name,
		SORTINGFIELD_PRICE: ListRequest_price,
		SORTINGFIELD_COUNT: ListRequest_changes_count,
		SORTINGFIELD_TIME:  ListRequest_timestamp,
	}
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
	Asc   int32
}

func ToPbFields(sort_field string) (ListRequest_SortingField, error) {
	val, ex := fields[sort_field]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}
