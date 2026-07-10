package main

import (
	"log"

	"github.com/lunadiotic/shopex-go/internal/config"
	"github.com/lunadiotic/shopex-go/internal/infrastructure/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// init logger
	logg := logger.New(cfg.Logger.Level)
	logg.Info("starting the application...")
	logg.Info(
		"configuration loaded",
		"app", cfg.App.Name,
		"version", cfg.App.Version,
		"environment", cfg.App.Env,
	)
}