package repository

import (
	"coop-gardens-be/internal/models"

	"gorm.io/gorm"
)

type CropRepository struct {
	DB *gorm.DB
}

func NewCropRepository(db *gorm.DB) *CropRepository {
	return &CropRepository{DB: db}
}

func (r *CropRepository) CreateCrop(crop *models.Crop) error {
	return r.DB.Create(crop).Error
}

func (r *CropRepository) GetAllCrops() ([]models.Crop, error) {
	var crops []models.Crop
	err := r.DB.Find(&crops).Error
	return crops, err
}

func (r *CropRepository) GetCropsBySeason(seasonID uint) ([]models.Crop, error) {
	var crops []models.Crop
	err := r.DB.Where("season_id = ?", seasonID).Find(&crops).Error
	return crops, err
}

func (r *CropRepository) AddCropToSeason(crop *models.Crop) error {
	return r.DB.Create(crop).Error
}
