package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/shopex-go/internal/delivery/http/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		response.Success(
			"Service is healthy", 
			gin.H{
				"status": "ok",
			},
		),
	)
}