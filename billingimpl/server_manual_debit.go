package billingimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leaq-ru/billing/safeerr"
	"github.com/leaq-ru/proto/codegen/go/billing"
	"github.com/leaq-ru/proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (s *server) ManualDebit(
	ctx context.Context,
	req *billing.ManualDebitRequest,
) (
	res *empty.Empty,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetUserId() == "" || req.GetAmount() == 0 {
		err = safeerr.BadRequest
		return
	}

	userOID, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	us, err := s.userClient.GetById(ctx, &user.GetByIdRequest{
		UserId: req.GetUserId(),
	})
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}
	if us.GetId() == "" {
		err = safeerr.BadRequest
		return
	}

	sess, err := s.mongoStartSession()
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	_, err = sess.WithTransaction(ctx, func(sc mongo.SessionContext) (_ interface{}, e error) {
		e = s.balanceModel.Inc(sc, userOID, req.GetAmount())
		if e != nil {
			return
		}

		e = s.invoiceModel.CreateSuccessDebitManual(sc, userOID, req.GetAmount())
		return
	})
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	res = &empty.Empty{}
	return
}
