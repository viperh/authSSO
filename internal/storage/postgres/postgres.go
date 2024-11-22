package postgres

import (
	"authSSO/internal/config"
	slog "authSSO/internal/log"
	"authSSO/internal/storage/postgres/perms"
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
	Perm *perms.Storage
}

func NewStorage(db *gorm.DB, log *slog.Logger, cfg *config.Config) *Storage {
	return &Storage{
		Db:     db,
		Config: cfg,

		Log:  log,
		User: user.NewStorage(db, log),
		Role: role.NewStorage(db, log),
		Perm: perms.NewStorage(db, log),
	}
}
