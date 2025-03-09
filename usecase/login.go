package usecase

import (
	"errors"
	_ "errors"
	_ "os"

	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"

	_ "golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// if err != nil {
	// 	return "", errors.New("Invalid password")
	// }

	if user.Password != password {
		return "", errors.New("Invalid password")
	}

	token, err := middlewares.GenerateToken(user.ID, user.Username, user.Email)

	if err != nil {
		return "", err
	}

	return token, nil
}
