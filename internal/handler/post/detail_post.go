package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) DetailPost(c *gin.Context) {
	ctx := c.Request.Context()
	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, statusCode, err := h.postService.DetailPost(ctx, postID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, result)
}
