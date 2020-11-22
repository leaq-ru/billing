package invoice

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (m Model) CreatePendingDebit(
	ctx context.Context,
	userOID primitive.ObjectID,
	rkInvoiceID uint64,
	amount uint32,
) (
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	count, err := m.invoices.CountDocuments(ctx, bson.M{
		"u": userOID,
		"s": 1,
		"o": 1,
		"ca": bson.M{
			"$gte": time.Now().UTC().AddDate(0, 0, -1),
		},
	}, options.Count().SetLimit(30))
	if err != nil {
		return
	}

	if count >= 30 {
		err = errors.New("create invoice limit exceeded. Try again later")
		return
	}

	_, err = m.invoices.InsertOne(ctx, Invoice{
		UserID:      userOID,
		RKInvoiceID: rkInvoiceID,
		Amount:      amount,
		Op:          Op_debit,
		Status:      Status_pending,
		CreatedAt:   time.Now().UTC(),
	})
	return
}
