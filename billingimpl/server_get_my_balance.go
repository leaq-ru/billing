package billingimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-billing/md"
	"github.com/nnqq/scr-billing/safeerr"
	"github.com/nnqq/scr-proto/codegen/go/billing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *server) GetMyBalance(
	ctx context.Context,
	_ *empty.Empty,
) (
	res *billing.GetMyBalanceResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	authUserID, err := md.GetUserID(ctx)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	authUserOID, err := primitive.ObjectIDFromHex(authUserID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	amount, err := s.balanceModel.Get(ctx, authUserOID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	res = &billing.GetMyBalanceResponse{
		Balance: amount,
	}
	return
}
