package event_log

import "go.mongodb.org/mongo-driver/mongo"

type Model struct {
	db        *mongo.Database
	eventLogs *mongo.Collection
}
