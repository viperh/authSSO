package controller

import "gorm.io/gorm"

var db *gorm.DB

func NewService(gorm *gorm.DB) {
	db = gorm
}
