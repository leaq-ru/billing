package robokassa

import (
	"github.com/leaq-ru/billing/balance"
	"github.com/leaq-ru/billing/event_log"
	"github.com/leaq-ru/billing/invoice"
	"github.com/nats-io/stan.go"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewWebhook(
	logger zerolog.Logger,
	stanConn stan.Conn,
	eventLogModel event_log.Model,
	balanceModel balance.Model,
	invoiceModel invoice.Model,
	mongoStartSession func(opts ...*options.SessionOptions) (mongo.Session, error),
	serviceName,
	webhookSecret,
	passwordTwo,
	isTest string,
) Webhook {
	return Webhook{
		logger:            logger,
		stanConn:          stanConn,
		eventLogModel:     eventLogModel,
		balanceModel:      balanceModel,
		invoiceModel:      invoiceModel,
		mongoStartSession: mongoStartSession,
		serviceName:       serviceName,
		webhookSecret:     webhookSecret,
		passwordTwo:       passwordTwo,
		isTest:            isTest,
		state: &state{
			done: make(chan struct{}),
		},
	}
}
