package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
)

func Signup(email, username, password string) error {
	user := &models.User{
		Email:    email,
		Username: username,
		Password: password,
	}

	err := repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
