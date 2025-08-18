package http

import (
	"context"
	"errors"
	"money/internal/app"
	"money/internal/config"
	v1Handlers "money/internal/http/handlers/v1"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Server is the struct that holds the echo Server
type Server struct {
	e *echo.Echo
}

// NewServer creates a new echo server
func NewServer() *Server {
	e := echo.New()

	e.HideBanner = true
	e.Server.ReadTimeout = config.Cfg.HTTPServer.ReadTimeout
	e.Server.WriteTimeout = config.Cfg.HTTPServer.WriteTimeout
	e.Server.ReadHeaderTimeout = config.Cfg.HTTPServer.ReadHeaderTimeout
	e.Server.IdleTimeout = config.Cfg.HTTPServer.IdleTimeout

	return &Server{
		e: e,
	}
}

// Serve starts the echo server and listens on the configured port and
// use middlewares for logging, metrics, tracing and locale
func (s *Server) Serve() {
	s.e.GET("/", v1Handlers.Index)

	//apiV1 := s.e.Group("/api/v1")

	go func() {
		<-app.A.Ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.e.Shutdown(ctx); err != nil {
			s.e.Logger.Fatal(err)
		}
	}()

	if err := s.e.Start(config.Cfg.HTTPServer.Listen); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.e.Logger.Fatal("shutting down the server")
	}
}
