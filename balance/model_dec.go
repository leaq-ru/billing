package balance

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) Dec(ctx context.Context, userID primitive.ObjectID, amount uint32) (ok bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := m.balances.UpdateOne(ctx, bson.M{
		"u": userID,
		"a": bson.M{
			"$gte": amount,
		},
	}, bson.M{
		"$inc": bson.M{
			"a": -int64(amount),
		},
	})
	if err != nil {
		return
	}

	ok = res.ModifiedCount == 1
	return
}
