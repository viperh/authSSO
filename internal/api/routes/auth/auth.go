package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(gin *gin.Engine) {
	authRoutes := gin.Group("/auth")
	{
		authRoutes.POST("/login")
		authRoutes.POST("/register")
	}
}
