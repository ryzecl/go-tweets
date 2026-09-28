package post

import (
	"go-tweets/internal/middleware"
	"go-tweets/internal/service/post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	api         *gin.Engine
	validate    *validator.Validate
	postService post.PostService
}

func NewPostHandler(api *gin.Engine, validate *validator.Validate, postService post.PostService) *PostHandler {
	return &PostHandler{
		api:         api,
		validate:    validate,
		postService: postService,
	}
}

func (h *PostHandler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/tweets")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/", h.CreatePost)
	routeAuth.PUT("/:post_id/update", h.UpdatePost)
	routeAuth.DELETE("/:post_id/delete", h.DeletePost)
	routeAuth.POST("/action", h.LikeOrUnlikePost)

	routeWithoutAuth := h.api.Group("/tweets")
	routeWithoutAuth.GET("/:post_id/detail", h.DetailPost)
	routeWithoutAuth.GET("/", h.GetAllPost)
}
