package invoice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (m Model) Get(
	ctx context.Context,
	userID primitive.ObjectID,
	skip,
	limit uint32,
) (
	res []Invoice,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cur, err := m.invoices.Find(ctx, Invoice{
		UserID: userID,
	}, options.Find().
		SetSort(bson.M{
			"_id": -1,
		}).
		SetSkip(int64(skip)).
		SetLimit(int64(limit)))
	if err != nil {
		return
	}

	err = cur.All(ctx, &res)
	return
}
