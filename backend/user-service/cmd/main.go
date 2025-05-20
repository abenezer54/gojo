package main

import (
	"fmt"
	"log"

	"github.com/abenezer54/gojo/backend/user-service/config"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("📦 Starting user-service...")
	config.InitDB()
	r := gin.Default()

	// Simple health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "user-service is running"})
	})

	log.Println("Starting user-service on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
