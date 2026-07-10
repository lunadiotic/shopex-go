package container

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/lunadiotic/shopex-go/internal/delivery/http/handler"
	"github.com/lunadiotic/shopex-go/internal/delivery/http/middleware"
	httpRouter "github.com/lunadiotic/shopex-go/internal/delivery/http/router"
)

type Container struct {
	Router *gin.Engine
}

func New(logger *slog.Logger) *Container {
	loggerMiddleware := middleware.Logger(logger)

	healthHandler := handler.NewHealthHandler()

	router := httpRouter.New(healthHandler, loggerMiddleware)

	return &Container{Router: router}
}