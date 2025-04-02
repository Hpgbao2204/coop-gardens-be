package models

import "time"

type Task struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `gorm:"not null" json:"title"`
    Description string    `json:"description"`
    Status      string    `gorm:"default:Pending" json:"status"`
    AssignedTo  string    `gorm:"type:uuid" json:"assigned_to"` // Change to string for UUID
    SeasonID    uint      `gorm:"index" json:"season_id"` 
    Season      Season    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
    Crops       []Crop    `gorm:"many2many:task_crops" json:"crops,omitempty"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}