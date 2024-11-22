package roles

import (
	"authSSO/internal/config"
	rlog "authSSO/internal/log"
	"authSSO/internal/models"
	"authSSO/internal/storage/postgres/role"
	"context"
)

type Service struct {
	db  *role.Storage
	cfg *config.Config
	log *rlog.Logger
}

func NewService(roleProvider *role.Storage, cfg *config.Config, log *rlog.Logger) *Service {
	return &Service{
		db:  roleProvider,
		cfg: cfg,
		log: log,
	}
}

func (s *Service) CreateRole(ctx context.Context, role *models.Role) error {
	result := s.db.Db.WithContext(ctx).Create(role)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) GetRoleByID(ctx context.Context, id uint) (*models.Role, error) {
	return s.GetRoleByID(ctx, id)
}

func (s *Service) UpdateRole(ctx context.Context, role *models.Role) error {
	return s.UpdateRole(ctx, role)
}

func (s *Service) DeleteRole(ctx context.Context, role *models.Role) error {
	return s.DeleteRole(ctx, role)
}
