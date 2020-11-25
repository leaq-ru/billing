package invoice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) CreateSuccessDebitRK(
	ctx context.Context,
	userOID primitive.ObjectID,
	rkInvoiceID uint64,
	amount uint32,
) (
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = m.invoices.InsertOne(ctx, Invoice{
		UserID:    userOID,
		Amount:    amount,
		Kind:      kind_debitRobokassa,
		Status:    status_success,
		CreatedAt: time.Now().UTC(),
		DebitRobokassa: &debitRobokassa{
			InvoiceID: rkInvoiceID,
		},
	})
	return
}
