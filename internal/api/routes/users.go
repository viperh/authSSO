package routes

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(gin *gin.Engine) {
	usersRoutes := gin.Group("/users")
	{
		usersRoutes.GET("/getAll")
		usersRoutes.POST("/create")
		usersRoutes.POST("/update")
		usersRoutes.POST("/delete")
	}
}
