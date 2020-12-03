package invoice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) CreateSuccessCreditDataPremium(
	ctx context.Context,
	userID primitive.ObjectID,
	amount,
	monthAmount uint32,
) (
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = m.invoices.InsertOne(ctx, Invoice{
		UserID:    userID,
		Amount:    amount,
		Kind:      kind_creditDataPremium,
		Status:    status_success,
		CreatedAt: time.Now().UTC(),
		CreditDataPremium: &creditDataPremium{
			MonthAmount: monthAmount,
		},
	})
	return
}
