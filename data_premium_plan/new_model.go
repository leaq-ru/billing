package data_premium_plan

import (
	"github.com/nnqq/scr-billing/mongo"
	m "go.mongodb.org/mongo-driver/mongo"
)

func NewModel(db *m.Database) Model {
	return Model{
		db:               db,
		dataPremiumPlans: db.Collection(mongo.CollDataPremiumPlan),
	}
}
