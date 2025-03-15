package usecase

import (
	"context"
	"errors"
	_ "errors"
	_ "os"

	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"

	_ "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, email, password string) (string, error) {
	ctx := context.Background()
	user, err := repository.NewUserRepository(db).GetUserByEmail(ctx, email)

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
