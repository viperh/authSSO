package users

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

func (s *Service) CreateUser(ctx context.Context, user *models.User) error {
	result := s.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	result := s.db.WithContext(ctx).Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	user := &models.User{}
	result := s.db.WithContext(ctx).Where("id = ?", id).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *models.User) error {
	result := s.db.WithContext(ctx).Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, user *models.User) error {
	result := s.db.WithContext(ctx).Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
