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

type state struct {
	sub               stan.Subscription
	subscribeCalledOK bool
	done              chan struct{}
}

type Webhook struct {
	logger            zerolog.Logger
	stanConn          stan.Conn
	eventLogModel     event_log.Model
	balanceModel      balance.Model
	invoiceModel      invoice.Model
	mongoStartSession func(opts ...*options.SessionOptions) (mongo.Session, error)
	serviceName       string
	webhookSecret     string
	passwordTwo       string
	isTest            string
	*state
}
