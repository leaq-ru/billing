package billingimpl

import (
	"github.com/nnqq/scr-billing/balance"
	"github.com/nnqq/scr-billing/counter"
	"github.com/nnqq/scr-billing/invoice"
	"github.com/nnqq/scr-billing/robokassa"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewServer(
	logger zerolog.Logger,
	invoiceModel invoice.Model,
	counterModel counter.Model,
	balanceModel balance.Model,
	companyClient parser.CompanyClient,
	userClient user.UserClient,
	robokassaClient robokassa.Client,
	robokassaWebhook robokassa.Webhook,
	mongoStartSession func(opts ...*options.SessionOptions) (mongo.Session, error),
) *server {
	return &server{
		logger:            logger,
		invoiceModel:      invoiceModel,
		counterModel:      counterModel,
		balanceModel:      balanceModel,
		companyClient:     companyClient,
		userClient:        userClient,
		robokassaClient:   robokassaClient,
		robokassaWebhook:  robokassaWebhook,
		mongoStartSession: mongoStartSession,
	}
}
