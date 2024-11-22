package perms

import (
	"authSSO/internal/config"
	rlog "authSSO/internal/log"
	"authSSO/internal/models"
	"context"
	"gorm.io/gorm"
)

type Storage struct {
	db  *gorm.DB
	cfg *config.Config
	log *rlog.Logger
}

func NewStorage(db *gorm.DB, log *rlog.Logger) *Storage {
	return &Storage{
		db:  db,
		log: log,
	}
}

func (s *Storage) GetPermsByUserId(ctx context.Context, userId uint64) ([]models.UserRole, error) {
	var roles []models.UserRole
	result := s.db.WithContext(ctx).Where("id = ?", userId).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (s *Storage) AddUserRoleById(ctx context.Context, userId uint64, roleId uint64) error {
	result := s.db.WithContext(ctx).Create(&models.UserRole{UserID: userId, RoleID: roleId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) DeleteUserRoleById(ctx context.Context, id uint64, id2 uint64) error {
	result := s.db.WithContext(ctx).Delete(&models.UserRole{}, "user_id = ? AND role_id = ?", id, id2)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
