package repository

import (
	"coop-gardens-be/config"
	"coop-gardens-be/internal/models"
)

func CreateUser(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}
