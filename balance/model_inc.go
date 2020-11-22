package balance

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (m Model) Inc(ctx context.Context, userID primitive.ObjectID, amount uint32) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = m.balances.UpdateOne(ctx, Balance{
		UserID: userID,
	}, bson.M{
		"$inc": Balance{
			Amount: amount,
		},
	}, options.Update().SetUpsert(true))
	return
}
