package billingimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leaq-ru/billing/balance"
	"github.com/leaq-ru/billing/md"
	"github.com/leaq-ru/billing/price"
	"github.com/leaq-ru/billing/safeerr"
	"github.com/leaq-ru/proto/codegen/go/billing"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (s *server) RenewDataPremium(
	ctx context.Context,
	req *billing.RenewDataPremiumRequest,
) (
	res *empty.Empty,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetMonthAmount() == 0 {
		err = errors.New("monthAmount required")
		return
	}

	authUserOID, err := md.GetUserOID(ctx)
	if err != nil {
		return
	}

	sess, err := s.mongoStartSession()
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	_, err = sess.WithTransaction(ctx, func(sc mongo.SessionContext) (_ interface{}, e error) {
		amount := req.GetMonthAmount() * price.DataPremium

		e = s.balanceModel.Dec(sc, authUserOID, amount)
		if e != nil {
			return
		}

		e = s.invoiceModel.CreateSuccessCreditDataPremium(sc, authUserOID, amount, req.GetMonthAmount())
		if e != nil {
			return
		}

		e = s.dataPremiumPlanModel.Renew(sc, authUserOID, req.GetMonthAmount())
		return
	})
	if err != nil {
		if errors.Is(err, balance.ErrInsufficientFunds) {
			return
		}

		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	res = &empty.Empty{}
	return
}
