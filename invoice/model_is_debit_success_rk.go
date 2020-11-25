package invoice

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m Model) IsDebitSuccessRK(ctx context.Context, rkInvoiceID uint64) (success bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = m.invoices.FindOne(ctx, bson.M{
		"k":    kind_debitRobokassa,
		"s":    status_success,
		"dr.i": rkInvoiceID,
	}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = nil
			return
		}
		return
	}

	success = true
	return
}
