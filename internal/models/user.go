package models

import (
	_ "gorm.io/gorm"

	"time"
)

type User struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email      string    `gorm:"uniqueIndex;not null" json:"email"`
	Password   string    `gorm:"not null" json:"-"`
	FullName   string    `gorm:"not null" json:"full_name"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"`
	GoogleID   *string   `json:"google_id,omitempty"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
