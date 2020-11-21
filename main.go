package main

import (
	"github.com/nnqq/scr-billing/billingimpl"
	"github.com/nnqq/scr-billing/call"
	"github.com/nnqq/scr-billing/config"
	"github.com/nnqq/scr-billing/counter"
	"github.com/nnqq/scr-billing/invoice"
	"github.com/nnqq/scr-billing/logger"
	"github.com/nnqq/scr-billing/mongo"
	"github.com/nnqq/scr-billing/robokassa"
	"github.com/nnqq/scr-billing/stan"
	graceful "github.com/nnqq/scr-lib-graceful"
	"github.com/nnqq/scr-proto/codegen/go/billing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"strings"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	companyService, err := call.New(cfg.Service.Parser)
	logg.Must(err)

	stanConn, err := stan.New(cfg.ServiceName, cfg.STAN.ClusterID, cfg.NATS.URL)
	logg.Must(err)

	db, err := mongo.New(cfg.ServiceName, cfg.MongoDB.URL)
	logg.Must(err)

	srv := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
	billing.RegisterBillingServer(srv, billingimpl.New(
		logg.ZL,
		invoice.New(db),
		counter.New(db),
		companyService,
		robokassa.New(
			cfg.Robokassa.WebhookSecret,
			cfg.Robokassa.MerchantLogin,
			cfg.Robokassa.PasswordOne,
			cfg.Robokassa.PasswordTwo,
			cfg.Robokassa.IsTest,
		),
	))

	go graceful.HandleSignals(srv.GracefulStop, func() {
		e := stanConn.Close()
		if e != nil {
			logg.ZL.Error().Err(e).Send()
		}
	})

	lis, err := net.Listen("tcp", strings.Join([]string{
		"0.0.0.0",
		cfg.Grpc.Port,
	}, ":"))
	logg.Must(err)
	logg.Must(srv.Serve(lis))
}
