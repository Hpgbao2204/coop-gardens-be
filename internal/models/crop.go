package models

import "time"

type Crop struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Type      string
    Status    string    `gorm:"default:Planted"`
    PlantedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}