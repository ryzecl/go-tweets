package user

import (
	"go-tweets/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Register(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RegisterRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, statusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, dto.RegisterResponse{
		ID: userID,
	})
}
