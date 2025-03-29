package repository

import (
	"coop-gardens-be/internal/models"

	"gorm.io/gorm"
)

type CropRepository struct {
	db *gorm.DB
}

func NewCropRepository(db *gorm.DB) *CropRepository {
	return &CropRepository{db}
}

func (r *CropRepository) CreateCrop(crop *models.Crop) error {
	return r.db.Create(crop).Error
}

func (r *CropRepository) GetAllCrops() ([]models.Crop, error) {
	var crops []models.Crop
	err := r.db.Find(&crops).Error
	return crops, err
}
