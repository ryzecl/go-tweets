package user

import (
	"go-tweets/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RefreshTokenRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.GetInt64("userID")
	token, refreshToken, statusCode, err := h.userService.RefreshToken(ctx, &req, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, dto.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}
