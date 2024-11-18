package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint64 `gorm:"not null"`
	RoleID uint64 `gorm:"not null"`
}
