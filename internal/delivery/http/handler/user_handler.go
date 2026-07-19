package handler

import (
	userDTO "github.com/lunadiotic/shopex-go/internal/delivery/http/dto/user"
	userUseCase "github.com/lunadiotic/shopex-go/internal/usecase/user"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase *userUseCase.UseCase
}

func NewUserHandler(
	useCase *userUseCase.UseCase,
) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req userDTO.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// err := h.useCase.Register(
	// 	c.Request.Context(),
	// 	req,
	// )

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}