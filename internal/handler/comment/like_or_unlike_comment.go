package comment

import (
	"go-tweets/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CommentHandler) LikeOrUnlikeComment(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.LikeOrUnlikeCommentRequest
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
	statusCode, err := h.commentService.LikeOrUnlikeComment(ctx, req.CommentID, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": "succesfully liked or unliked comment",
	})
}
