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

// GetByID lấy thông tin cây trồng theo ID
func (r *CropRepository) GetByID(id uint) (*models.Crop, error) {
	var crop models.Crop
	err := r.DB.First(&crop, id).Error
	if err != nil {
		return nil, err
	}
	return &crop, nil
}

// Update cập nhật thông tin cây trồng
func (r *CropRepository) Update(crop *models.Crop) error {
	return r.DB.Save(crop).Error
}
