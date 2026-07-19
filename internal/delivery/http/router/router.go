package router

import (
	"github.com/gin-gonic/gin"

	"github.com/lunadiotic/shopex-go/internal/delivery/http/handler"
)

func New(
	healthHandler *handler.HealthHandler, 
	userHandler *handler.UserHandler,
	middleware ...gin.HandlerFunc,
) *gin.Engine {
	r := gin.New()
	r.Use(middleware...)

	r.GET("/health", healthHandler.Check)
	r.Group("/api/v1/auth")
	{
		r.POST("/register", userHandler.Register)
	}

	return r
}