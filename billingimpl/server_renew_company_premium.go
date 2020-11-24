package billingimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-billing/md"
	"github.com/nnqq/scr-billing/premium"
	"github.com/nnqq/scr-proto/codegen/go/billing"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (s *server) RenewCompanyPremium(
	ctx context.Context,
	req *billing.RenewCompanyPremiumRequest,
) (
	res *empty.Empty,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetCompanyId() == "" || req.GetMonthAmount() == 0 {
		err = badRequest
		return
	}

	authUserID, err := md.GetUserID(ctx)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	authUserOID, err := primitive.ObjectIDFromHex(authUserID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = internalServerError
		return
	}

	comp, err := s.companyClient.GetBy(ctx, &parser.GetByRequest{
		CompanyId: req.GetCompanyId(),
	})
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = internalServerError
		return
	}
	if comp.GetId() == "" {
		s.logger.Error().Err(err).Send()
		err = badRequest
		return
	}

	companyOID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = internalServerError
		return
	}

	sess, err := s.mongoStartSession()
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = internalServerError
		return
	}

	errInsufficientFunds := errors.New("insufficient funds")

	var renewSuccess bool
	_, err = sess.WithTransaction(ctx, func(sc mongo.SessionContext) (_ interface{}, e error) {
		amount := req.GetMonthAmount() * premium.MonthPrice

		ok, e := s.balanceModel.Dec(sc, authUserOID, amount)
		if e != nil {
			return
		}
		if !ok {
			e = errInsufficientFunds
			return
		}

		e = s.invoiceModel.CreateSuccessCredit(sc, authUserOID, companyOID, amount, req.GetMonthAmount())
		if e != nil {
			return
		}

		if !renewSuccess {
			_, e = s.companyClient.RenewCompanyPremium(sc, &parser.RenewCompanyPremiumRequest{
				CompanyId:   req.GetCompanyId(),
				MonthAmount: req.GetMonthAmount(),
			})
			if e != nil {
				return
			}

			renewSuccess = true
		}
		return
	})
	if err != nil {
		s.logger.Error().Err(err).Send()

		if errors.Is(err, errInsufficientFunds) {
			return
		}

		err = internalServerError
		return
	}

	res = &empty.Empty{}
	return
}
