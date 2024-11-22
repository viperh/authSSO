package perms

import (
	"authSSO/internal/config"
	rlog "authSSO/internal/log"
	"authSSO/internal/models"
	"authSSO/internal/storage/postgres/perms"
	"context"
)

type Service struct {
	db  *perms.Storage
	cfg *config.Config
	log *rlog.Logger
}

func NewService(permsProvider *perms.Storage, cfg *config.Config, log *rlog.Logger) *Service {
	return &Service{
		db:  permsProvider,
		cfg: cfg,
		log: log,
	}
}

func (s *Service) GetPermsByUserId(ctx context.Context, userId uint64) ([]models.UserRole, error) {
	return s.db.GetPermsByUserId(ctx, userId)
}

func (s *Service) AddUserRoleById(ctx context.Context, userId uint64, roleId uint64) error {
	return s.db.AddUserRoleById(ctx, userId, roleId)
}

func (s *Service) DeleteUserRoleById(ctx context.Context, userId uint64, roleId uint64) error {
	return s.db.DeleteUserRoleById(ctx, userId, roleId)
}
