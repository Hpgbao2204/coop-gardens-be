package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type CropUsecase struct {
	repo *repository.CropRepository
}

func NewCropUsecase(repo *repository.CropRepository) *CropUsecase {
	return &CropUsecase{repo}
}

func (u *CropUsecase) CreateCrop(crop *models.Crop) error {
	return u.repo.CreateCrop(crop)
}

func (u *CropUsecase) GetAllCrops() ([]models.Crop, error) {
	return u.repo.GetAllCrops()
}

func (u *CropUsecase) GetCropsBySeason(seasonID uint) ([]models.Crop, error) {
	return u.repo.GetCropsBySeason(seasonID)
}

func (u *CropUsecase) AddCropToSeason(seasonID uint, crop *models.Crop) error {
	// Kiểm tra xem mùa vụ có tồn tại không
	var season models.Season
	if err := u.repo.DB.First(&season, seasonID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("season not found")
		}
		return err
	}

	// Gán seasonID cho crop
	crop.SeasonID = seasonID

	// Lưu crop vào database
	return u.repo.AddCropToSeason(crop)
}
