package users

import (
	"authSSO/internal/config"
	rlog "authSSO/internal/log"
	"authSSO/internal/models"
	"authSSO/internal/storage"
	"authSSO/internal/storage/postgres/user"
	"context"
)

type Service struct {
	db  *user.Storage
	cfg *config.Config
	log *rlog.Logger
}

func NewService(userProvider *user.Storage, cfg *config.Config, log *rlog.Logger) *Service {
	return &Service{db: userProvider, cfg: cfg, log: log}
}

func (s *Service) CreateUser(ctx context.Context, email, username, password string) error {
	_, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return storage.ErrUserExists
	}
	return nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (s *Service) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	return nil, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, user *models.User) error {
	return nil
}
