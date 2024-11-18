package storage

import (
	"authSSO/internal/config"
	slog "authSSO/internal/log"
	"authSSO/internal/storage/postgres/role"
	"authSSO/internal/storage/postgres/user"
	"gorm.io/gorm"
)

type Storage struct {
	Db     *gorm.DB
	Config *config.Config

	Log  *slog.Logger
	User *user.Storage
	Role *role.Storage
}
