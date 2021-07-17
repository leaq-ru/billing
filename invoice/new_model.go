package invoice

import (
	"github.com/leaq-ru/billing/mongo"
	m "go.mongodb.org/mongo-driver/mongo"
)

func NewModel(db *m.Database) Model {
	return Model{
		db:       db,
		invoices: db.Collection(mongo.CollInvoice),
	}
}
