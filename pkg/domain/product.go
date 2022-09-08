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
