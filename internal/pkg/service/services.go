package service

import (
	"authSSO/internal/config"
	rlog "authSSO/internal/log"
	"authSSO/internal/pkg/service/auth"
	"authSSO/internal/pkg/service/perms"
	"authSSO/internal/pkg/service/roles"
	"authSSO/internal/pkg/service/users"
	"authSSO/internal/storage/postgres"
)

type Services struct {
	AuthService  *auth.Service
	UserService  *users.Service
	RolesService *roles.Service
	PermsService *perms.Service
}

func NewServices(providers *postgres.Storage, cfg *config.Config, log *rlog.Logger) *Services {
	db := providers.Db
	userProvider := providers.User
	roleProvider := providers.Role

	return &Services{
		AuthService:  auth.NewService(db),
		UserService:  users.NewService(userProvider, cfg, log),
		RolesService: roles.NewService(roleProvider, cfg, log),
		PermsService: perms.NewService(db),
	}
	
}
