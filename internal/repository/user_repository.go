package repository

import (
	"coop-gardens-be/config"
	"coop-gardens-be/internal/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
