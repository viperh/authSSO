package user

import (
	slog "authSSO/internal/log"
	"authSSO/internal/models"
	"authSSO/internal/storage"
	"context"
	"errors"
	"gorm.io/gorm"
)

type StorageInterface interface {
	CreateUser(ctx context.Context, email, password, username string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserById(ctx context.Context, id uint64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, user *models.User) error
}

type Storage struct {
	StorageInterface
	Db  *gorm.DB
	Log *slog.Logger
}

func NewStorage(db *gorm.DB, log *slog.Logger) *Storage {
	return &Storage{
		Db:  db,
		Log: log,
	}
}

func (f *Storage) CreateUser(ctx context.Context, email, password, username string) (*models.User, error) {

	user := models.User{
		Email:    email,
		Password: password,
		Username: username,
	}

	if err := f.Db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (f *Storage) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	var user models.User

	if err := f.Db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrUserNotExists
		}
		return nil, err
	}

	return &user, nil

}

func (f *Storage) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	var user models.User

	if err := f.Db.WithContext(ctx).First(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrUserNotExists
		}
		return nil, err
	}

	return &user, nil
}

func (f *Storage) GetUserById(ctx context.Context, id uint64) (*models.User, error) {

	var user models.User

	if err := f.Db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrUserNotExists
		}
		return nil, err
	}

	return &user, nil
}

func (f *Storage) UpdateUser(ctx context.Context, user *models.User) error {

	if err := f.Db.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (f *Storage) DeleteUser(ctx context.Context, user *models.User) error {

	if err := f.Db.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
