package invoice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) CreateSuccessDebit(
	ctx context.Context,
	userOID primitive.ObjectID,
	rkInvoiceID uint64,
	amount uint32,
) (
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = m.invoices.InsertOne(ctx, invoice{
		UserID:      userOID,
		RKInvoiceID: rkInvoiceID,
		Amount:      amount,
		Op:          Op_debit,
		Status:      Status_success,
		CreatedAt:   time.Now().UTC(),
	})
	return
}
