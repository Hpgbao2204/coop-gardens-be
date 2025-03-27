package models

import (
	_ "gorm.io/gorm"
)

// Role định nghĩa quyền
type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

// UserRole liên kết User và Role (N-N)
type UserRole struct {
	UserID string `gorm:"primaryKey; type:uuid"`
	RoleID uint   `gorm:"primaryKey"`
}
