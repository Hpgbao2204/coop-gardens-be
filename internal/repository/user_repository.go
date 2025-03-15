package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"coop-gardens-be/config"
	"coop-gardens-be/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	HandleForgotPassword(ctx context.Context, email string) error
	UpdateUserPassword(ctx context.Context, userId int64, currentPassword string, newPassword string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// GetUserByEmail is a function to get user by email
func (u *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	// Set a default timeout for operations
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var user models.User
	result := config.DB.WithContext(ctx).Where("email = ?", email).First(&user)
	// Perform query
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("User with email %s not found", email)
			return nil, result.Error
		}
		if errors.Is(result.Error, context.DeadlineExceeded) {
			log.Printf("Query timed out for email %s", email)
			return nil, result.Error
		}
		return nil, result.Error
	}

	return &user, nil
}

// HandleForgotPassword is a function to handle forgot password
func (u *UserRepositoryImpl) HandleForgotPassword(ctx context.Context, email string) error {
	// Set a default timeout for operations
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user models.User

	// Query user by email
	result := u.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("No user found with email: %s", email)
			return fmt.Errorf("no user found with email: %s", email)
		}
		log.Printf("Failed to fetch user with Email: %s, error: %v", email, result.Error)
		return result.Error
	}

	// Generate a new password
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(90000000) + 10000000

	// Encrypt the random number to use as the new password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", randomNumber)), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}

	// Update the user's password in the database
	user.Password = string(encryptedPassword)
	if err := u.db.WithContext(ctx).Save(&user).Error; err != nil {
		log.Printf("Failed to update password for email: %s, error: %v", email, err)
		return nil, nil, err
	}

	// Return the generated random password as a string
	resultStr := fmt.Sprintf("%d", randomNumber)
	return &user.Username, &resultStr, nil
}

// UpdateUserPassword implements UserRepository.
func (u *UserRepositoryImpl) UpdateUserPassword(ctx context.Context, userId int64, currentPassword string, newPassword string) error {
	// Set a timeout for the operation
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel() // Ensure resources are released

	// Define a variable to hold the result
	var user models.User

	// Query the database
	result := u.db.WithContext(ctx).First(&user, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("No user found with ID: %d", userId)
			return result.Error
		}
		log.Printf("Failed to fetch user with ID: %d, error: %v", userId, result.Error)
		return result.Error
	}

	// Compare the current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword)); err != nil {
		return errors.New("invalid current password")
	}

	// Encrypt the new password
	encryptedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash new password: %v", err)
		return err
	}

	// Update the user's password in the database
	user.Password = string(encryptedNewPassword)
	if err := u.db.WithContext(ctx).Save(&user).Error; err != nil {
		log.Printf("Failed to update password for user ID: %d, error: %v", userId, err)
		return err
	}

	return nil
}

// DeleteUser implements UserRepository.
func (u *UserRepositoryImpl) DeleteUser(ctx context.Context, userId int64) error {
	// Set a timeout for the operation
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Delete the user by ID
	result := u.db.WithContext(ctx).Delete(&models.User{}, userId)
	if result.Error != nil {
		log.Printf("Failed to delete user with ID: %d, error: %v", userId, result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Printf("No user found with ID: %d", userId)
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	// Set a timeout for the operation
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Create the user
	result := u.db.WithContext(ctx).Create(user) // field in models.User
	if result.Error != nil {
		log.Printf("Failed to create user, error: %v", result.Error)
		return result.Error
	}

	return nil
}
