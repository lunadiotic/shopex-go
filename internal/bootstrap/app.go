package bootstrap

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/shopex-go/internal/config"
	httpserver "github.com/lunadiotic/shopex-go/internal/delivery/http/server"
	"github.com/lunadiotic/shopex-go/internal/infrastructure/logger"
)

type Application struct {
	config *config.Config
	logger *slog.Logger
	router *gin.Engine
	server *http.Server
}

func New() (*Application, error) {
	// load config
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// init logger
	logg, err := logger.New(cfg.Logger.Level)
	if err != nil {
		return nil, err
	}
	
	// init router
	router := gin.New()
	srv := httpserver.New(cfg.Server, router)

	// init application
	app := &Application{
		config: cfg,
		logger: logg,
		router: router,
		server: srv,
	}

	return app, nil
}

func (a *Application) Run() error {
	a.logger.Info(
		"HTTP server starting",
		"address",
		a.server.Addr,
	)

	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("HTTP server failed", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	
	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	a.logger.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), a.config.Server.ShutdownTimeout)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error("HTTP server shutdown failed", "error", err)
		return err
	}

	a.logger.Info("Server stopped gracefully")

	return nil
}