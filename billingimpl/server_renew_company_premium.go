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
	"github.com/leaq-ru/proto/codegen/go/parser"
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
		err = safeerr.BadRequest
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
		err = safeerr.InternalServerError
		return
	}

	comp, err := s.companyClient.GetBy(ctx, &parser.GetByRequest{
		CompanyId: req.GetCompanyId(),
	})
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}
	if comp.GetId() == "" {
		s.logger.Error().Err(err).Send()
		err = safeerr.BadRequest
		return
	}

	companyOID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	sess, err := s.mongoStartSession()
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	var renewSuccess bool
	_, err = sess.WithTransaction(ctx, func(sc mongo.SessionContext) (_ interface{}, e error) {
		amount := req.GetMonthAmount() * price.CompanyPremium

		e = s.balanceModel.Dec(sc, authUserOID, amount)
		if e != nil {
			return
		}

		e = s.invoiceModel.CreateSuccessCreditCompanyPremium(sc, authUserOID, companyOID, amount, req.GetMonthAmount())
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
