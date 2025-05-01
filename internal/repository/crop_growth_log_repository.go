package repository

import (
	"coop-gardens-be/internal/models"

	"gorm.io/gorm"
)

type CropGrowthLogRepository struct {
	DB *gorm.DB
}

func NewCropGrowthLogRepository(db *gorm.DB) *CropGrowthLogRepository {
	return &CropGrowthLogRepository{DB: db}
}

// Create adds a new crop growth log to the database
func (r *CropGrowthLogRepository) Create(log *models.CropGrowthLog) error {
	return r.DB.Create(log).Error
}

// GetByCropID retrieves all growth logs for a specific crop
func (r *CropGrowthLogRepository) GetByCropID(cropID uint) ([]models.CropGrowthLog, error) {
	var logs []models.CropGrowthLog
	result := r.DB.Where("crop_id = ?", cropID).Order("log_date desc").Find(&logs)
	return logs, result.Error
}

// GetByID retrieves a specific growth log by its ID
func (r *CropGrowthLogRepository) GetByID(id uint) (*models.CropGrowthLog, error) {
	var log models.CropGrowthLog
	result := r.DB.First(&log, id)
	return &log, result.Error
}

// Update modifies an existing crop growth log
func (r *CropGrowthLogRepository) Update(log *models.CropGrowthLog) error {
	return r.DB.Save(log).Error
}

// Delete removes a crop growth log from the database
func (r *CropGrowthLogRepository) Delete(id uint) error {
	return r.DB.Delete(&models.CropGrowthLog{}, id).Error
}
