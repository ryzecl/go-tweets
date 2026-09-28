package post

import (
	"go-tweets/internal/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")

	page, _ := strconv.ParseInt(pageStr, 10, 64)
	limit, _ := strconv.ParseInt(limitStr, 10, 64)

	param := dto.GetAllPostRequest{
		Page:  page,
		Limit: limit,
	}

	result, statusCode, err := h.postService.GetAllPost(ctx, &param)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, result)
}
