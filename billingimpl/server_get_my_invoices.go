package billingimpl

import (
	"context"
	"github.com/leaq-ru/billing/md"
	"github.com/leaq-ru/billing/pagination"
	"github.com/leaq-ru/billing/safeerr"
	"github.com/leaq-ru/proto/codegen/go/billing"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *server) GetMyInvoices(
	ctx context.Context,
	req *billing.GetMyInvoicesRequest,
) (
	res *billing.GetMyInvoicesResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit, err := pagination.ApplyDefaultLimit(req)
	if err != nil {
		return
	}

	authUserID, err := md.GetUserID(ctx)
	if err != nil {
		return
	}

	authUserOID, err := primitive.ObjectIDFromHex(authUserID)
	if err != nil {
		err = safeerr.BadRequest
		return
	}

	invoices, err := s.invoiceModel.Get(ctx, authUserOID, req.GetOpts().GetSkip(), limit)
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = safeerr.InternalServerError
		return
	}

	var companyIDs []string
	for _, inv := range invoices {
		if inv.CreditCompanyPremium != nil {
			companyIDs = append(companyIDs, inv.CreditCompanyPremium.CompanyID.Hex())
		}
	}

	type (
		compID   = string
		compSlug = string
	)
	slugs := map[compID]compSlug{}
	if len(companyIDs) != 0 {
		resComps, e := s.companyClient.GetManyBy(ctx, &parser.GetManyByRequest{
			CompanyIds: companyIDs,
		})
		if e != nil {
			s.logger.Error().Err(err).Send()
			err = safeerr.InternalServerError
			return
		}

		for _, comp := range resComps.GetCompanies() {
			slugs[comp.GetId()] = comp.GetSlug()
		}
	}

	res = &billing.GetMyInvoicesResponse{}
	for _, inv := range invoices {
		resInv := &billing.InvoiceItem{
			Id:        inv.ID.Hex(),
			CreatedAt: inv.CreatedAt.String(),
			Amount:    inv.Amount,
			Status:    billing.Status(inv.Status),
			Kind:      billing.Kind(inv.Kind),
		}
		if inv.DebitRobokassa != nil {
			resInv.DebitRobokassa = &billing.DebitRobokassa{
				InvoiceId: inv.DebitRobokassa.InvoiceID,
			}
		}
		if inv.CreditCompanyPremium != nil {
			cID := inv.CreditCompanyPremium.CompanyID.Hex()
			slug, ok := slugs[cID]
			if !ok {
				s.logger.Error().
					Str("companyID", cID).
					Msg("expected to get company slug but nothing found")
				err = safeerr.InternalServerError
				return
			}

			resInv.CreditCompanyPremium = &billing.CreditCompanyPremium{
				CompanyId:   cID,
				MonthAmount: inv.CreditCompanyPremium.MonthAmount,
				CompanySlug: slug,
			}
		}
		if inv.CreditDataPremium != nil {
			resInv.CreditDataPremium = &billing.CreditDataPremium{
				MonthAmount: inv.CreditDataPremium.MonthAmount,
			}
		}

		res.Invoices = append(res.Invoices, resInv)
	}
	return
}
