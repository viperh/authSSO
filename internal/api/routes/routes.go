package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(gin *gin.Engine) {
	RegisterAuthRoutes(gin)
	RegisterPermissionsRoutes(gin)
	RegisterRolesRoutes(gin)
	RegisterUserRoutes(gin)
}
