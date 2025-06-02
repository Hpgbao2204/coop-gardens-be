package repository

import (
	_ "coop-gardens-be/config"
	"coop-gardens-be/internal/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"errors"
	"log"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
	// Make sure we're not losing the password during creation
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	// Log to help debug
	log.Printf("User created with ID: %s and email: %s", user.ID, user.Email)

	return nil
}

func (r *UserRepository) CheckUserExists(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
    if email == "" {
        return nil, errors.New("email is required")
    }
    var user models.User
    result := r.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("Roles").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserRoles(userID string) ([]models.Role, error) {
	var roles []models.Role
	if err := r.DB.Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *UserRepository) AssignRoleToUser(userID string, roleName string) error {
	// First, find the role ID by name
	var role models.Role
	if err := r.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
		return err
	}

	// Create the user-role relationship
	userRole := models.UserRole{
		UserID: userID, // No need to convert, directly use the UUID string
		RoleID: role.ID,
	}

	return r.DB.Create(&userRole).Error
}

func (r *UserRepository) HasRole(userID string, roleName string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Role{}).
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ? AND roles.name = ?", userID, roleName).
		Count(&count).Error

	return count > 0, err
}
