package perms

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

func (s *Service) GetPermsByUserId(ctx context.Context, userId uint64) ([]models.Role, error) {
	var roles []models.Role
	result := s.db.WithContext(ctx).Where("id = ?", userId).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (s *Service) AddUserRoleById(ctx context.Context, userId uint64, roleId uint64) error {
	result := s.db.WithContext(ctx).Create(&models.UserRole{UserID: userId, RoleID: roleId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) DeleteUserRoleById(ctx context.Context, userId uint64, roleId uint64) error {
	result := s.db.WithContext(ctx).Where("user_id = ? AND role_id = ?", userId, roleId).Delete(&models.UserRole{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
