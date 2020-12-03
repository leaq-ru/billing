package data_premium_plan

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type dataPremiumPlan struct {
	ID              primitive.ObjectID `bson:"_id"`
	UserID          primitive.ObjectID `bson:"u"`
	PremiumDeadline time.Time          `bson:"pd"`
}
