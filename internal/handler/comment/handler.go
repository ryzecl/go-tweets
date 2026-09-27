package comment

import (
	"go-tweets/internal/middleware"
	"go-tweets/internal/service/comment"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentHandler struct {
	api            *gin.Engine
	validate       *validator.Validate
	commentService comment.CommentService
}

func NewCommentHandler(api *gin.Engine, validate *validator.Validate, commentService comment.CommentService) *CommentHandler {
	return &CommentHandler{
		api:            api,
		validate:       validate,
		commentService: commentService,
	}
}

func (h *CommentHandler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/comment")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/", h.CreateComment)
}
