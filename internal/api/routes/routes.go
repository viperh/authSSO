package routes

import (
	"authSSO/internal/api/routes/auth"
	"authSSO/internal/api/routes/permissions"
	"authSSO/internal/api/routes/roles"
	"authSSO/internal/api/routes/users"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(gin *gin.Engine) {
	auth.RegisterAuthRoutes(gin)
	permissions.RegisterPermissionsRoutes(gin)
	roles.RegisterRolesRoutes(gin)
	users.RegisterUserRoutes(gin)
}
