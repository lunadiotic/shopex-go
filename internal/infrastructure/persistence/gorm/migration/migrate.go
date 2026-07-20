package migration

import (
	"log/slog"

	userModel "github.com/lunadiotic/shopex-go/internal/infrastructure/persistence/gorm/user"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, logger *slog.Logger) error {
	logger.Info("Database migration started")

	defer logger.Info("Database migration completed")

	return db.AutoMigrate(&userModel.Model{})
}