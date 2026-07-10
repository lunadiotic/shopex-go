package bootstrap

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/shopex-go/internal/config"
	"github.com/lunadiotic/shopex-go/internal/delivery/http/server"
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
	srv := server.New(cfg.Server, router)

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
	a.logger.Info("starting the application...", "address", a.server.Addr)

	return a.server.ListenAndServe()
}