package main

import (
	"fmt"
	"go-tweets/internal/config"
	userHandler "go-tweets/internal/handler/user"
	userRepo "go-tweets/internal/repository/user"
	userService "go-tweets/internal/service/user"
	"go-tweets/pkg/internalsql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mesage": "App is running",
		})
	})

	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(cfg, userRepo)
	userHandler := userHandler.NewUserHandler(r, userService)
	userHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
