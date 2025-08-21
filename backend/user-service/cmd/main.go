package main

import (
	"log"

	"github.com/abenezer54/gojo/backend/user-service/config"
	"github.com/abenezer54/gojo/backend/user-service/internal/controller"
	"github.com/abenezer54/gojo/backend/user-service/internal/repository"
	"github.com/abenezer54/gojo/backend/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	repo := repository.NewUserRepository(config.DB)
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)

	// Router
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/signup", controller.SignUp)
	}

	// Start server
	log.Println("Server starting on port 8080")
	r.Run(":8080")
}
