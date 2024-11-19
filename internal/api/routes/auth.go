package routes

import (
	"authSSO/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(gin *gin.Engine) {
	authRoutes := gin.Group("/auth")
	{
		authRoutes.POST("/login", controller.Login)
		authRoutes.POST("/register")
	}
}
