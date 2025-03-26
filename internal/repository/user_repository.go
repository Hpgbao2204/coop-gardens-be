package repository

import (
	_ "coop-gardens-be/config"
	"coop-gardens-be/internal/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) CheckUserExists(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
