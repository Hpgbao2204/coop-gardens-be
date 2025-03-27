package usecase

import (
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	UserRepo *repository.UserRepository
}

// Xử lý đăng ký user
func (u *AuthUsecase) Signup(user *models.User) error {
	exists, _ := u.UserRepo.CheckUserExists(user.Email)
	if exists {
		return errors.New("email đã được sử dụng")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save user to database
	if err := u.UserRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

// Xử lý đăng nhập
func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("email hoặc mật khẩu không đúng")
	}

	// For initial plaintext password only
	if user.Password == password {
		token, err := middlewares.GenerateJWT(user.ID)
		if err != nil {
			return "", errors.New("cannot generate token")
		}
		return token, nil
	}

	// Kiểm tra mật khẩu
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("email hoặc mật khẩu không đúng")
	}

	// Tạo JWT
	token, err := middlewares.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("không thể tạo token")
	}

	return token, nil
}
