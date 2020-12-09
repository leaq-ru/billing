package data_premium_plan

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (m Model) Renew(ctx context.Context, userID primitive.ObjectID, monthAmount uint32) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	month := time.Duration(monthAmount) * 31 * 24 * time.Hour

	update, err := m.dataPremiumPlans.UpdateOne(ctx, dataPremiumPlan{
		UserID: userID,
	}, bson.M{
		"$setOnInsert": dataPremiumPlan{
			PremiumDeadline: time.Now().Add(month).UTC(),
		},
	}, options.Update().SetUpsert(true))
	if err != nil {
		return
	}

	if update.UpsertedCount == 1 {
		return
	}

	_, err = m.dataPremiumPlans.UpdateOne(ctx, dataPremiumPlan{
		UserID: userID,
	}, bson.A{bson.M{
		"$set": bson.M{
			"pd": bson.M{
				"$add": bson.A{
					"$pd",
					month.Milliseconds(),
				},
			},
		},
	}})
	return
}
