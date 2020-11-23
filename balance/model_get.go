package balance

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m Model) Get(ctx context.Context, userID primitive.ObjectID) (amount uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var doc balance
	err = m.balances.FindOne(ctx, balance{
		UserID: userID,
	}).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = nil
			return
		}

		return
	}

	amount = doc.Amount
	return
}
