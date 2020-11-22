package billingimpl

import (
	"github.com/nnqq/scr-billing/counter"
	"github.com/nnqq/scr-billing/invoice"
	"github.com/nnqq/scr-billing/robokassa"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/rs/zerolog"
)

func NewServer(
	logger zerolog.Logger,
	invoiceModel invoice.Model,
	counterModel counter.Model,
	companyClient parser.CompanyClient,
	robokassaClient robokassa.Client,
	robokassaWebhook robokassa.Webhook,
) *server {
	return &server{
		logger:           logger,
		invoiceModel:     invoiceModel,
		counterModel:     counterModel,
		companyClient:    companyClient,
		robokassaClient:  robokassaClient,
		robokassaWebhook: robokassaWebhook,
	}
}
