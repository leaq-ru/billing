package invoice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) CreateSuccessCredit(
	ctx context.Context,
	userOID primitive.ObjectID,
	companyID primitive.ObjectID,
	amount,
	monthAmount uint32,
) (
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = m.invoices.InsertOne(ctx, Invoice{
		UserID:      userOID,
		CompanyID:   companyID,
		MonthAmount: monthAmount,
		Amount:      amount,
		Op:          Op_credit,
		Status:      Status_success,
		CreatedAt:   time.Now().UTC(),
	})
	return
}
