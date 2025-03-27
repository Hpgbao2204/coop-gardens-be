package usecase

import (
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	UserRepo *repository.UserRepository
}

func (u *AuthUsecase) Signup(user *models.User, roles []string) error {
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

	// Assign roles if provided
	if len(roles) > 0 {
		for _, roleName := range roles {
			if err := u.UserRepo.AssignRoleToUser(user.ID, roleName); err != nil {
				// Log error but continue
				fmt.Printf("Error assigning role %s: %v\n", roleName, err)
			}
		}
	} else {
		// Assign default "User" role
		if err := u.UserRepo.AssignRoleToUser(user.ID, "User"); err != nil {
			return err
		}
	}

	return nil
}

func (uc *AuthUsecase) SignupWithRole(user *models.User, role string) error {
	// Check if user already exists
	exists, err := uc.UserRepo.CheckUserExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save user to database
	if err := uc.UserRepo.CreateUser(user); err != nil {
		return err
	}

	// Assign role to user
	if err := uc.UserRepo.AssignRoleToUser(user.ID, role); err != nil {
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
