package billingimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-billing/md"
	"github.com/nnqq/scr-billing/price"
	"github.com/nnqq/scr-billing/robokassa"
	"github.com/nnqq/scr-billing/safeerr"
	"github.com/nnqq/scr-proto/codegen/go/billing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func (s *server) CreateInvoice(ctx context.Context, req *billing.CreateInvoiceRequest) (
	res *billing.CreateInvoiceResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetAmount() < price.CompanyPremium {
		err = errors.New("amount too small")
		return
	}

	authUserID, err := md.GetUserID(ctx)
	if err != nil {
		return
	}

	authUserOID, err := primitive.ObjectIDFromHex(authUserID)
	if err != nil {
		err = errors.New(http.StatusText(http.StatusBadRequest))
		return
	}

	rkInvoiceID, err := s.counterModel.GetNextRKInvoiceID(ctx)
	if err != nil {
		s.logger.Error().
			Str("userID", authUserID).
			Uint32("amount", req.GetAmount()).
			Err(err).
			Send()
		err = safeerr.InternalServerError
		return
	}

	err = s.invoiceModel.CreatePendingDebitRK(ctx, authUserOID, rkInvoiceID, req.GetAmount())
	if err != nil {
		s.logger.Error().
			Str("userID", authUserID).
			Uint32("amount", req.GetAmount()).
			Err(err).
			Send()
		err = safeerr.InternalServerError
		return
	}

	paymentURL, err := s.robokassaClient.CreatePaymentURL(rkInvoiceID, req.GetAmount())
	if err != nil {
		if errors.Is(err, robokassa.ErrOnlyIntegerRub) {
			return
		}

		s.logger.Error().
			Str("userID", authUserID).
			Uint32("amount", req.GetAmount()).
			Err(err).
			Send()
		err = safeerr.InternalServerError
		return
	}

	res = &billing.CreateInvoiceResponse{
		Url: paymentURL,
	}
	return
}
