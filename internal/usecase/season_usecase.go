package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
)

type SeasonUsecase struct {
	repo *repository.SeasonRepository
}

func NewSeasonUsecase(repo *repository.SeasonRepository) *SeasonUsecase {
	return &SeasonUsecase{repo}
}

func (u *SeasonUsecase) CreateSeason(season *models.Season) error {
	// Validate season
	if season.Name == "" {
		return ErrValidation{Message: "Season name is required"}
	}

	// Make sure end date is after start date
	if !season.EndDate.After(season.StartDate) {
		return ErrValidation{Message: "End date must be after start date"}
	}

	return u.repo.CreateSeason(season)
}

func (u *SeasonUsecase) GetAllSeasons() ([]models.Season, error) {
	return u.repo.GetAllSeasons()
}

func (u *SeasonUsecase) GetSeasonByID(id uint) (*models.Season, error) {
	return u.repo.GetSeasonByID(id)
}

func (u *SeasonUsecase) GetSeasonWithCrops(id uint) (*models.Season, error) {
	return u.repo.GetSeasonWithCrops(id)
}

func (u *SeasonUsecase) UpdateSeason(season *models.Season) error {
	return u.repo.UpdateSeason(season)
}

func (u *SeasonUsecase) DeleteSeason(id uint) error {
	return u.repo.DeleteSeason(id)
}

// Custom error type for validation
type ErrValidation struct {
	Message string
}

func (e ErrValidation) Error() string {
	return e.Message
}
