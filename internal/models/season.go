package models

import "time"

type Season struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Status    string    `gorm:"default:Planning"`
	Crops     []Crop    `gorm:"foreignKey:SeasonID"` // Quan hệ 1:N với Crops
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
