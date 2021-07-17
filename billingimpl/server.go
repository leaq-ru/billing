package billingimpl

import (
	"github.com/leaq-ru/billing/balance"
	"github.com/leaq-ru/billing/counter"
	"github.com/leaq-ru/billing/data_premium_plan"
	"github.com/leaq-ru/billing/invoice"
	"github.com/leaq-ru/billing/robokassa"
	"github.com/leaq-ru/proto/codegen/go/billing"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"github.com/leaq-ru/proto/codegen/go/user"
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
