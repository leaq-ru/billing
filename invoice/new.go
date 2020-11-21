package invoice

import (
	"github.com/nnqq/scr-billing/mongo"
	m "go.mongodb.org/mongo-driver/mongo"
)

func New(db *m.Database) Model {
	return Model{
		db:       db,
		invoices: db.Collection(mongo.CollInvoice),
		balances: db.Collection(mongo.CollBalance),
	}
}
