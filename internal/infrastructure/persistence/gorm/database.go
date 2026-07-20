package gorm

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/lunadiotic/shopex-go/internal/config"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg config.DatabaseConfig, logger *slog.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
		cfg.Timezone,
	)

	db, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	if err := sqlDB.Ping(); err != nil {
		logger.Error("Database connection failed", "error", err)
		return nil, err
	}

	logger.Info("Database connection successful")

	return db, nil
}