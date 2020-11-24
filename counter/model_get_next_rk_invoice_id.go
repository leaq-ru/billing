package counter

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (m Model) GetNextRKInvoiceID(ctx context.Context) (seq uint64, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var doc counter
	err = m.counters.FindOneAndUpdate(ctx, counter{
		Entity: robokassaInvoiceID,
	}, bson.M{
		"$inc": bson.M{
			"s": 1,
		},
	}, options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After),
	).Decode(&doc)
	if err != nil {
		return
	}

	seq = doc.Sequence
	return
}
