package usecase

import (
    "coop-gardens-be/internal/models"
    "coop-gardens-be/internal/repository"
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