package counter

import (
	"github.com/nnqq/scr-billing/mongo"
	m "go.mongodb.org/mongo-driver/mongo"
)

func New(db *m.Database) Model {
	return Model{
		db:       db,
		counters: db.Collection(mongo.CollCounter),
	}
}
