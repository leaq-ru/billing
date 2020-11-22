package invoice

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Model struct {
	db       *mongo.Database
	invoices *mongo.Collection
}
