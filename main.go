package main

import (
	"github.com/nnqq/scr-billing/balance"
	"github.com/nnqq/scr-billing/billingimpl"
	"github.com/nnqq/scr-billing/call"
	"github.com/nnqq/scr-billing/config"
	"github.com/nnqq/scr-billing/counter"
	"github.com/nnqq/scr-billing/data_premium_plan"
	"github.com/nnqq/scr-billing/event_log"
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
	"sync"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	companyClient, userClient, err := call.NewClients(cfg.Service.Parser, cfg.Service.User)
	logg.Must(err)

	stanConn, err := stan.NewConn(cfg.ServiceName, cfg.STAN.ClusterID, cfg.NATS.URL)
	logg.Must(err)

	db, err := mongo.NewConn(cfg.ServiceName, cfg.MongoDB.URL)
	logg.Must(err)

	rkWebhook := robokassa.NewWebhook(
		logg.ZL,
		stanConn,
		event_log.NewModel(db),
		balance.NewModel(db),
		invoice.NewModel(db),
		db.Client().StartSession,
		cfg.ServiceName,
		cfg.Robokassa.WebhookSecret,
		cfg.Robokassa.PasswordTwo,
		cfg.Robokassa.IsTest,
	)
	logg.Must(rkWebhook.Subscribe())

	grpcSrv := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcSrv, health.NewServer())
	billing.RegisterBillingServer(grpcSrv, billingimpl.NewServer(
		logg.ZL,
		invoice.NewModel(db),
		counter.NewModel(db),
		balance.NewModel(db),
		data_premium_plan.NewModel(db),
		companyClient,
		userClient,
		robokassa.NewClient(
			cfg.Robokassa.MerchantLogin,
			cfg.Robokassa.PasswordOne,
			cfg.Robokassa.IsTest,
		),
		rkWebhook,
		db.Client().StartSession,
	))

	lis, err := net.Listen("tcp", strings.Join([]string{
		"0.0.0.0",
		cfg.Grpc.Port,
	}, ":"))
	logg.Must(err)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		graceful.HandleSignals(grpcSrv.GracefulStop, rkWebhook.GracefulStop)
	}()
	go func() {
		defer wg.Done()
		logg.Must(grpcSrv.Serve(lis))
	}()
	go func() {
		defer wg.Done()
		logg.Must(rkWebhook.Serve())
	}()
	wg.Wait()
}
