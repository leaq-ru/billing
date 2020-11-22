package robokassa

import (
	"github.com/nats-io/stan.go"
	"github.com/nnqq/scr-billing/balance"
	"github.com/nnqq/scr-billing/event_log"
	"github.com/nnqq/scr-billing/invoice"
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
	passwordTwo string,
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
		state: &state{
			done: make(chan struct{}),
		},
	}
}
