package usecase

import (
	"errors"

	"coop-gardens-be/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func Signup(email, username, password string) error {
	existingUser, _ := repository.GetUserByEmail(email)
	if existingUser.ID != 0 {
		return errors.New("Email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	err = repository.CreateUser(email, username, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
