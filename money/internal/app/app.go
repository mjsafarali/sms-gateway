package app

import (
	"api-gateway/internal/config"
	"api-gateway/internal/repositories"
	"api-gateway/internal/services"
	"api-gateway/log"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	goredis "github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// application is the main application struct that holds all the dependencies
type application struct {
	RedisClient *goredis.Client
	DB          *sqlx.DB
	Ctx         context.Context
	cancelFunc  context.CancelFunc
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

// WithDatabase initializes the database connection
func WithDatabase() {
	cfg := config.Cfg.Database

	db, err := sqlx.Open("mysql", cfg.DSN())
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(cfg.MaxConn)
	db.SetMaxIdleConns(cfg.IdleConn)

	ticker := time.NewTicker(cfg.DialTimeout)
	connected := false
	connectionAttempt := 0
	for !connected && connectionAttempt < cfg.DialRetry {
		err := db.PingContext(context.Background())
		if err == nil {
			log.Info("Connection established successfully to database")
			connected = true
			ticker.Stop()
			break
		}

		select {
		case <-A.Ctx.Done():
			connectionAttempt = cfg.DialRetry
		case <-ticker.C:
			err := db.PingContext(context.Background())
			if err == nil {
				log.Info("Connection established successfully to database")
				connected = true
				ticker.Stop()
			} else {
				log.Info("Database connection failed. Attempting to connect again. err: %s", err.Error())
				connectionAttempt++
			}
		}
	}

	ticker.Stop()

	if !connected {
		log.Fatal("Failed to connect to Database")
	}

	A.DB = db
}

// WithRedis initializes the redis client
func WithRedis() {
	cfg := config.Cfg.Redis
	opt := &goredis.Options{
		Addr:        cfg.Host + ":" + cfg.Port,
		DialTimeout: cfg.DialTimeout,
		Username:    cfg.Username,
		Password:    cfg.Password,
	}

	A.RedisClient = goredis.NewClient(opt)
}

// WithRepositories initializes the repositories
func WithRepositories() {
	repositories.Transactions = repositories.NewMysqlTransaction(A.DB)
}

func WithServices() {
	services.WalletServiceInstance = services.NewWalletService(repositories.Transactions)
}
