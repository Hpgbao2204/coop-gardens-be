package models

import "time"

type Crop struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Type      string
	SeasonID  uint      `gorm:"index"`
	Season    Season    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status    string    `gorm:"default:Planted"`
	PlantedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
