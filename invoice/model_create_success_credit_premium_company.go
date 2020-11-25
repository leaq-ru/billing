package invoice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) CreateSuccessCreditPremiumCompany(
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
		UserID:    userOID,
		Amount:    amount,
		Kind:      kind_creditCompanyPremium,
		Status:    status_success,
		CreatedAt: time.Now().UTC(),
		CreditCompanyPremium: &creditCompanyPremium{
			CompanyID:   companyID,
			MonthAmount: monthAmount,
		},
	})
	return
}
