package data_premium_plan

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m Model) GetByUserID(
	ctx context.Context,
	userID primitive.ObjectID,
) (
	premium bool,
	premiumDeadline time.Time,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var doc dataPremiumPlan
	err = m.dataPremiumPlans.FindOne(ctx, dataPremiumPlan{
		UserID: userID,
	}).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = nil
		}
		return
	}

	premium = true
	premiumDeadline = doc.PremiumDeadline
	return
}
