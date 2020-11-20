package invoice

import "go.mongodb.org/mongo-driver/mongo"

const Coll = "invoices"

type Model struct {
	db       *mongo.Database
	invoices *mongo.Collection
}
