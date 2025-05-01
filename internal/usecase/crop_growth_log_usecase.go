package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
	"fmt"
	// "time" // Removed unused import
)

type CropGrowthLogUsecase struct {
	CropGrowthLogRepo *repository.CropGrowthLogRepository
	CropRepo          *repository.CropRepository // Needed to update Crop's GrowthStage
}

func NewCropGrowthLogUsecase(logRepo *repository.CropGrowthLogRepository, cropRepo *repository.CropRepository) *CropGrowthLogUsecase {
	return &CropGrowthLogUsecase{
		CropGrowthLogRepo: logRepo,
		CropRepo:          cropRepo,
	}
}

// CreateLog adds a new crop growth log and updates the crop's current growth stage
func (uc *CropGrowthLogUsecase) CreateLog(log *models.CropGrowthLog) error {
	// Set LogDate if not provided
	// if log.LogDate.IsZero() {
	// 	log.LogDate = time.Now()
	// }

	// Create the log entry
	err := uc.CropGrowthLogRepo.Create(log)
	if err != nil {
		return fmt.Errorf("failed to create growth log: %w", err)
	}

	// Update the corresponding Crop's GrowthStage
	crop, err := uc.CropRepo.GetByID(log.CropID)
	if err != nil {
		// Log the error but don't fail the whole operation, as the log was created
		fmt.Printf("Warning: Failed to retrieve crop %d after creating growth log: %v\n", log.CropID, err)
		return nil // Or return a specific warning error
	}

	if crop.GrowthStage != log.GrowthStage {
		crop.GrowthStage = log.GrowthStage
		err = uc.CropRepo.Update(crop)
		if err != nil {
			// Log the error but don't fail the whole operation
			fmt.Printf("Warning: Failed to update crop %d growth stage after creating growth log: %v\n", log.CropID, err)
		}
	}

	return nil
}

// GetLogsByCropID retrieves all growth logs for a specific crop
func (uc *CropGrowthLogUsecase) GetLogsByCropID(cropID uint) ([]models.CropGrowthLog, error) {
	return uc.CropGrowthLogRepo.GetByCropID(cropID)
}

// GetLogByID retrieves a specific growth log by its ID
func (uc *CropGrowthLogUsecase) GetLogByID(id uint) (*models.CropGrowthLog, error) {
	return uc.CropGrowthLogRepo.GetByID(id)
}

// UpdateLog modifies an existing crop growth log
// Consider if updating the Crop's GrowthStage is needed here as well
func (uc *CropGrowthLogUsecase) UpdateLog(log *models.CropGrowthLog) error {
	// Potentially add logic to update Crop.GrowthStage if this log is the latest one
	return uc.CropGrowthLogRepo.Update(log)
}

// DeleteLog removes a crop growth log
func (uc *CropGrowthLogUsecase) DeleteLog(id uint) error {
	return uc.CropGrowthLogRepo.Delete(id)
}
