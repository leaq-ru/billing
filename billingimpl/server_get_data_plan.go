package billingimpl

import (
	"context"
	"github.com/nnqq/scr-proto/codegen/go/billing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *server) GetDataPlan(
	ctx context.Context,
	req *billing.GetDataPlanRequest,
) (
	res *billing.GetDataPlanResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	userOID, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	premium, premiumDeadline, err := s.dataPremiumPlanModel.GetByUserID(ctx, userOID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	res = &billing.GetDataPlanResponse{
		Premium:         premium,
		PremiumDeadline: premiumDeadline.String(),
	}
	return
}
