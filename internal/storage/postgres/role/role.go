package role

import (
	slog "authSSO/internal/log"
	"authSSO/internal/models"
	"authSSO/internal/storage"
	"context"
	"errors"
	"gorm.io/gorm"
)

type roleStorage interface {
	CreateRole(ctx context.Context, name, description string) (*models.Role, error)
	GetRoleById(ctx context.Context, roleId uint64) (*models.Role, error)
	GetRoleByName(ctx context.Context, name string) (*models.Role, error)
	UpdateRole(ctx context.Context, roleId uint64, name, description string) error
	DeleteRole(ctx context.Context, roleId uint64) error
}

type Storage struct {
	roleStorage
	Db  *gorm.DB
	Log *slog.Logger
}

func NewStorage(db *gorm.DB, log *slog.Logger) *Storage {
	return &Storage{
		Db:  db,
		Log: log,
	}
}
func (f *Storage) CreateRole(ctx context.Context, name, description string) (*models.Role, error) {

	_, err := f.GetRoleByName(ctx, name)
	if err == nil {
		return nil, storage.ErrRoleExists
	}

	role := models.Role{
		Name:        name,
		Description: description,
	}

	if err := f.Db.WithContext(ctx).Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil

}

func (f *Storage) GetRoleById(ctx context.Context, roleId uint64) (*models.Role, error) {
	var role models.Role

	if err := f.Db.WithContext(ctx).First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrRoleNotExists
		}
		return nil, err
	}

	return &role, nil

}

func (f *Storage) GetRoleByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role

	if err := f.Db.WithContext(ctx).First(&role, name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrRoleNotExists
		}
		return nil, err
	}

	return &role, nil
}

func (f *Storage) UpdateRole(ctx context.Context, roleId uint64, name, description string) error {

	role := models.Role{
		Name:        name,
		Description: description,
	}

	if err := f.Db.WithContext(ctx).Model(&models.Role{}).Where("id = ?", roleId).Updates(&role).Error; err != nil {
		return err
	}

	return nil
}

func (f *Storage) DeleteRole(ctx context.Context, roleId uint64) error {
	if err := f.Db.WithContext(ctx).Where("id = ?", roleId).Delete(&models.Role{}).Error; err != nil {
		return err
	}

	return nil
}
