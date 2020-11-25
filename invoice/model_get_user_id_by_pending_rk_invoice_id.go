package invoice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) GetUserIDByPendingRKInvoiceID(
	ctx context.Context,
	rkInvoiceID uint64,
) (
	userID primitive.ObjectID,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var doc Invoice
	err = m.invoices.FindOne(ctx, bson.M{
		"k":    kind_debitRobokassa,
		"s":    status_pending,
		"dr.i": rkInvoiceID,
	}).Decode(&doc)
	if err != nil {
		return
	}

	userID = doc.UserID
	return
}
