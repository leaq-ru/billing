package invoice

import "go.mongodb.org/mongo-driver/mongo"

func New(db *mongo.Database) Model {
	return Model{
		db:       db,
		invoices: db.Collection(Coll),
	}
}
