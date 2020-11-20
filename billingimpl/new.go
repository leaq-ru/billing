package billingimpl

import (
	"github.com/nnqq/scr-billing/invoice"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/rs/zerolog"
)

func New(logger zerolog.Logger, invoiceModel invoice.Model, companyClient parser.CompanyClient) *server {
	return &server{
		logger:        logger,
		invoiceModel:  invoiceModel,
		companyClient: companyClient,
	}
}
