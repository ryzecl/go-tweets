package user

import (
	"go-tweets/internal/service/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	api         *gin.Engine
	userService user.UserService
}

func NewUserHandler(api *gin.Engine, userService user.UserService) *UserHandler {
	return &UserHandler{
		api:         api,
		userService: userService,
	}
}

func (h *UserHandler) RouteList() {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
}
