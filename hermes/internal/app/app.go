package app

import (
	"context"
	"hermes/internal/config"
	"hermes/internal/services"
	"hermes/log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nats-io/nats.go"
)

// application is the main application struct that holds all the dependencies
type application struct {
	Ctx        context.Context
	cancelFunc context.CancelFunc
	NatsJS     nats.JetStreamContext
}

var (
	// A is the singleton instance of application
	A *application
)

func init() {
	A = &application{}
}

// WithGracefulShutdown registers a signal handler for graceful shutdown
func WithGracefulShutdown() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	A.Ctx, A.cancelFunc = context.WithCancel(context.Background())

	go func() {
		sig := <-c
		log.Info("system call", sig)
		A.cancelFunc()
	}()
}

// WithNats initializes the NATS connection
func WithNats() {
	cfg := config.Cfg.Nats
	opts := []nats.Option{
		nats.ReconnectWait(cfg.ConnectWait),
		nats.Timeout(cfg.DialTimeout),
		nats.FlusherTimeout(cfg.FlusherTimeout),
		nats.PingInterval(cfg.PingInterval),
		nats.ReconnectBufSize(cfg.ConnectBufSize),
		nats.SyncQueueLen(cfg.MaxChanLen),
		nats.MaxPingsOutstanding(cfg.MaxPingOut),
	}

	natsConn, err := nats.Connect(cfg.Address, opts...)
	if err != nil {
		log.Fatalf("error in nats Connect, err: %+v", err.Error())
	}

	js, err := natsConn.JetStream()
	if err != nil {
		log.Fatalf("error in creating jetstream, err: %+v", err.Error())
	}

	A.NatsJS = js
	log.Info("Connection established successfully to jetstream")
}

func WithServices() {
	services.SenderSrv = services.NewLogSender()
	services.PublisherSrv = services.NewNatsPublisher(A.NatsJS)
}
