package billingimpl

import (
	"github.com/nnqq/scr-billing/balance"
	"github.com/nnqq/scr-billing/counter"
	"github.com/nnqq/scr-billing/data_premium_plan"
	"github.com/nnqq/scr-billing/invoice"
	"github.com/nnqq/scr-billing/robokassa"
	"github.com/nnqq/scr-proto/codegen/go/billing"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type server struct {
	billing.UnimplementedBillingServer
	logger               zerolog.Logger
	invoiceModel         invoice.Model
	counterModel         counter.Model
	balanceModel         balance.Model
	dataPremiumPlanModel data_premium_plan.Model
	companyClient        parser.CompanyClient
	userClient           user.UserClient
	robokassaClient      robokassa.Client
	robokassaWebhook     robokassa.Webhook
	mongoStartSession    func(opts ...*options.SessionOptions) (mongo.Session, error)
}
