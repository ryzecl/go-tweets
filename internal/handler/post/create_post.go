package post

import (
	"go-tweets/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreatePost(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.CreatePostRequest
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
	postID, statusCode, err := h.postService.CreatePost(ctx, &req, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, dto.CreatePostResponse{
		ID: postID,
	})
}
