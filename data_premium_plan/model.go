package data_premium_plan

import "go.mongodb.org/mongo-driver/mongo"

type Model struct {
	db               *mongo.Database
	dataPremiumPlans *mongo.Collection
}
