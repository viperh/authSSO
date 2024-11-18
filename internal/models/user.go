package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string  `gorm:"unique, not null"`
	Password string  `gorm:"not null"`
	Username string  `gorm:"unique, not null"`
	Roles    []*Role `gorm:"many2many:user_roles; constraint:OnDelete:CASCADE;"`
}
