package models

import "time"

type CropGrowthLog struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CropID       uint      `gorm:"index;not null" json:"crop_id"`
	Crop         Crop      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	LogDate      time.Time `gorm:"not null" json:"log_date"`
	GrowthStage  string    `json:"growth_stage"`  // e.g., Seedling, Vegetative, Flowering, Fruiting
	Height       float64   `json:"height"`        // in cm or other unit
	HealthStatus string    `json:"health_status"` // e.g., Healthy, Pest Infestation, Nutrient Deficiency
	Notes        string    `json:"notes"`
	ImageURL     string    `json:"image_url"` // Optional image of the plant at this stage
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
