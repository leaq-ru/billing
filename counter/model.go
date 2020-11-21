package counter

import "go.mongodb.org/mongo-driver/mongo"

type Model struct {
	db       *mongo.Database
	counters *mongo.Collection
}
