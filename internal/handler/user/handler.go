package user

import (
	"go-tweets/internal/middleware"
	"go-tweets/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	api         *gin.Engine
	validate    *validator.Validate
	userService user.UserService
}

func NewUserHandler(api *gin.Engine, validate *validator.Validate, userService user.UserService) *UserHandler {
	return &UserHandler{
		api:         api,
		validate:    validate,
		userService: userService,
	}
}

func (h *UserHandler) RouteList(secretKey string) {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
	authRoute.POST("/login", h.Login)

	refreshRoute := h.api.Group("/auth")
	refreshRoute.Use(middleware.AuthRefreshTokenMiddleware(secretKey))
	refreshRoute.POST("/refresh", h.RefreshToken)
}
