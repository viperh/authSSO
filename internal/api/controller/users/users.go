package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewService(gorm *gorm.DB) {
	db = gorm
}

func CreateUser(c *gin.Context) {

}
