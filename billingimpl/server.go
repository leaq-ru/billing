package billingimpl

import (
	"github.com/nnqq/scr-billing/invoice"
	"github.com/nnqq/scr-proto/codegen/go/billing"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/rs/zerolog"
)

type server struct {
	billing.UnimplementedBillingServer
	logger        zerolog.Logger
	invoiceModel  invoice.Model
	companyClient parser.CompanyClient
}
