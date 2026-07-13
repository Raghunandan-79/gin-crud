package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raghunandan/gin-crud/internal/config"
	"github.com/raghunandan/gin-crud/internal/models"
	"github.com/raghunandan/gin-crud/internal/routes"
)

func main() {
	// connecting database
	config.ConnectDB()

	// Automigrating the databse
	config.DB.AutoMigrate(&models.User{})

	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}