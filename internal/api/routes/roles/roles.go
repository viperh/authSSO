package roles

import "github.com/gin-gonic/gin"

func RegisterRolesRoutes(gin *gin.Engine) {
	rolesRoutes := gin.Group("/roles")
	{
		rolesRoutes.GET("/getAll")
		rolesRoutes.POST("/create")
		rolesRoutes.POST("/update")
		rolesRoutes.POST("/delete")
	}
}
