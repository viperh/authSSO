package service

import (
	"authSSO/internal/pkg/service/auth"
	"authSSO/internal/pkg/service/perms"
	"authSSO/internal/pkg/service/roles"
	"authSSO/internal/pkg/service/users"
	"gorm.io/gorm"
)

type Services struct {
	AuthService  *auth.Service
	UserService  *users.Service
	RolesService *roles.Service
	PermsService *perms.Service
}

func NewServices(db *gorm.DB) *Services {
	return &Services{
		AuthService:  auth.NewService(db),
		UserService:  users.NewService(db),
		RolesService: roles.NewService(db),
		PermsService: perms.NewService(db),
	}
}
