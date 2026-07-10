package router

import (
	"github.com/gin-gonic/gin"

	"github.com/lunadiotic/shopex-go/internal/delivery/http/handler"
)

func New() *gin.Engine {
	r := gin.New()

	healthHandler := handler.NewHealthHandler()
	r.GET("/health", healthHandler.Check)

	return r
}