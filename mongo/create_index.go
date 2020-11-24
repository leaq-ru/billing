package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndex(db *mongo.Database) (err error) {
	ctx := context.Background()

	_, err = db.Collection(CollInvoice).Indexes().CreateMany(ctx, []mongo.IndexModel{{
		Keys: bson.M{
			"r": 1,
		},
		Options: options.Index().SetPartialFilterExpression(bson.M{
			"r": bson.M{
				"$exists": true,
			},
		}),
	}, {
		Keys: bson.M{
			"u": 1,
		},
	}})
	if err != nil {
		return
	}

	_, err = db.Collection(CollBalance).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{
			"u": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	return
}
