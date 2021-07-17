package billingimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leaq-ru/billing/md"
	"github.com/leaq-ru/billing/safeerr"
	"github.com/leaq-ru/proto/codegen/go/billing"
	"time"
)

func (s *server) GetMyDataPlan(ctx context.Context, _ *empty.Empty) (res *billing.GetMyDataPlanResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	authUserOID, err := md.GetUserOID(ctx)
	if err != nil {
		return
	}

	premium, premiumDeadline, err := s.dataPremiumPlanModel.GetByUserID(ctx, authUserOID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	res = &billing.GetMyDataPlanResponse{
		Premium:         premium,
		PremiumDeadline: premiumDeadline.String(),
	}
	return
}
