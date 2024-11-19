package auth

import (
	"authSSO/internal/api/controller"
	"context"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Login(c context.Context, data controller.LoginRequest) {}

func (s *Service) Register() {

}
