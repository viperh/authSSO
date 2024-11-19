package routes

import "github.com/gin-gonic/gin"

func RegisterPermissionsRoutes(gin *gin.Engine) {
	permissionsRoutes := gin.Group("/permissions")
	{
		permissionsRoutes.GET("/getByUserId")
		permissionsRoutes.POST("/addRole")
		permissionsRoutes.POST("/removeRole")

	}
}
