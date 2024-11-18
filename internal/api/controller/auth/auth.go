package auth

import (
	"authSSO/internal/lib/bcrypt"
	"authSSO/internal/models"
	"authSSO/internal/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewService(gorm *gorm.DB) {
	db = gorm
}

type LoginRequest struct {
	Email    string
	Password string
}

func Login(c *gin.Context, req LoginRequest) (*models.User, error) {
	var user models.User

	result := db.WithContext(c).Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrUserNotExists
		}
	}

	if bcrypt.CheckPassword(req.Password, user.Password) {
		return &user, nil
	}

	return nil, storage.ErrAuth
}

type RegisterRequest struct {
	Email    string
	Password string
	Username string
}

func Register(c *gin.Context, req RegisterRequest) (*models.User, error) {

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	var searchedUser *models.User

	result := db.WithContext(c).Where("email = ?", req.Email).First(&searchedUser)
	if result.Error == nil {
		return nil, storage.ErrUserExists
	}

	result = db.WithContext(c).Create(&user)
	if result.Error != nil {
		return nil, errors.New("there was a problem while creating the user")
	}

	return user, nil
}
