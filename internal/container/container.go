package container

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	gormPersistence "github.com/lunadiotic/shopex-go/internal/infrastructure/persistence/gorm"
	"github.com/lunadiotic/shopex-go/internal/infrastructure/persistence/gorm/migration"
	gormUserRepository "github.com/lunadiotic/shopex-go/internal/infrastructure/persistence/gorm/user"
	bcryptHasher "github.com/lunadiotic/shopex-go/internal/infrastructure/security/bcrypt"

	"github.com/lunadiotic/shopex-go/internal/config"
	"github.com/lunadiotic/shopex-go/internal/delivery/http/handler"
	"github.com/lunadiotic/shopex-go/internal/delivery/http/middleware"
	httpRouter "github.com/lunadiotic/shopex-go/internal/delivery/http/router"
	domainUser "github.com/lunadiotic/shopex-go/internal/domain/user"
	userUseCase "github.com/lunadiotic/shopex-go/internal/usecase/user"
)

type Container struct {
	Router *gin.Engine
	DB     *gorm.DB
	UserRepository domainUser.Repository
	UserUseCase *userUseCase.UseCase
}

func New(cfg *config.Config, logger *slog.Logger) (*Container, error) {
	loggerMiddleware := middleware.Logger(logger)

	healthHandler := handler.NewHealthHandler()
	userHandler := handler.NewUserHandler(&userUseCase.UseCase{},)

	router := httpRouter.New(healthHandler, userHandler, loggerMiddleware)

	db, err := gormPersistence.NewDatabase(cfg.Database, logger)
	if err != nil {
		return nil, err
	}

	err = migration.AutoMigrate(db, logger)
	if err != nil {
		return nil, err
	}

	userRepository := gormUserRepository.NewRepository(db)
	hasher := bcryptHasher.NewHasher()
	userUseCase := userUseCase.NewUseCase(userRepository, hasher)

	return &Container{
		Router: router, 
		DB: db, 
		UserRepository: userRepository, 
		UserUseCase: userUseCase,
	} , nil
}