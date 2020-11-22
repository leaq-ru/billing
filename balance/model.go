package balance

import "go.mongodb.org/mongo-driver/mongo"

type Model struct {
	db       *mongo.Database
	balances *mongo.Collection
}
