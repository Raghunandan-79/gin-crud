package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raghunandan/gin-crud/internal/handlers"
	"github.com/raghunandan/gin-crud/internal/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)
	router.POST("/profile", middleware.AuthMiddleware, handlers.Profile)

	router.GET("/users", middleware.AuthMiddleware, handlers.GetAllUsers)
}