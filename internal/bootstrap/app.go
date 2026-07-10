package bootstrap

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/shopex-go/internal/config"
	"github.com/lunadiotic/shopex-go/internal/infrastructure/logger"
)

type Application struct {
	config *config.Config
	logger *slog.Logger
	router *gin.Engine
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

	// init application
	app := &Application{
		config: cfg,
		logger: logg,
		router: router,
	}

	return app, nil
}

func (a *Application) Run() error {
	a.logger.Info("starting the application...")

	return a.router.Run(
		fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port),
	)
}