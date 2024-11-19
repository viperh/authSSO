package roles

import (
	"authSSO/internal/models"
	"context"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetRoles(ctx context.Context) ([]*models.Role, error) {
	var roles []*models.Role
	result := s.db.WithContext(ctx).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}

func (s *Service) CreateRole(ctx context.Context, role *models.Role) error {
	result := s.db.WithContext(ctx).Create(role)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) GetRoleByID(ctx context.Context, id uint) (*models.Role, error) {
	var role models.Role
	result := s.db.WithContext(ctx).First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (s *Service) UpdateRole(ctx context.Context, role *models.Role) error {
	result := s.db.WithContext(ctx).Save(role)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) DeleteRole(ctx context.Context, role *models.Role) error {
	result := s.db.WithContext(ctx).Delete(role)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
