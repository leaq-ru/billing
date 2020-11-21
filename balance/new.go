package balance

import (
	"github.com/nnqq/scr-billing/mongo"
	m "go.mongodb.org/mongo-driver/mongo"
)

func New(db *m.Database) Model {
	return Model{
		db:       db,
		balances: db.Collection(mongo.CollBalance),
		invoices: db.Collection(mongo.CollInvoice),
	}
}
