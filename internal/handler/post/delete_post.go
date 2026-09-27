package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) DeletePost(c *gin.Context) {
	var (
		ctx       = c.Request.Context()
		userID    = c.GetInt64("userID")
		postIDStr = c.Param("post_id")
	)

	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	statusCode, err := h.postService.DeletePost(ctx, postID, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": "Post deleted successfully",
	})
}
