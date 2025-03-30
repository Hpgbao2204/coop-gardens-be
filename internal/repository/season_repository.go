package repository

import (
	"coop-gardens-be/internal/models"

	"gorm.io/gorm"
)

type SeasonRepository struct {
	DB *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *SeasonRepository {
	return &SeasonRepository{DB: db}
}

func (r *SeasonRepository) CreateSeason(season *models.Season) error {
	return r.DB.Create(season).Error
}

func (r *SeasonRepository) GetAllSeasons() ([]models.Season, error) {
	var seasons []models.Season
	err := r.DB.Find(&seasons).Error
	return seasons, err
}

func (r *SeasonRepository) GetSeasonByID(id uint) (*models.Season, error) {
	var season models.Season
	err := r.DB.First(&season, id).Error
	if err != nil {
		return nil, err
	}
	return &season, nil
}

func (r *SeasonRepository) GetSeasonWithCrops(id uint) (*models.Season, error) {
	var season models.Season
	err := r.DB.Preload("Crops").First(&season, id).Error
	if err != nil {
		return nil, err
	}
	return &season, nil
}

func (r *SeasonRepository) UpdateSeason(season *models.Season) error {
	return r.DB.Save(season).Error
}

func (r *SeasonRepository) DeleteSeason(id uint) error {
	return r.DB.Delete(&models.Season{}, id).Error
}
