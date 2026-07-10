package router

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/lunadiotic/shopex-go/internal/delivery/http/handler"
	"github.com/lunadiotic/shopex-go/internal/delivery/http/middleware"
)

func New(logger *slog.Logger) *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.Logger(logger),
	)

	healthHandler := handler.NewHealthHandler()
	r.GET("/health", healthHandler.Check)

	return r
}