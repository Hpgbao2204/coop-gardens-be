package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey;autoIncrement"` // User ID
	Email    string `gorm:"not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	// FirstName 		string `gorm:"not null"`
	// LastName 			string `gorm:"not null"`
	// Gender        string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
