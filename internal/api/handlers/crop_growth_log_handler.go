package handlers

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CropGrowthLogHandler struct {
	CropGrowthLogUC *usecase.CropGrowthLogUsecase
}

func NewCropGrowthLogHandler(uc *usecase.CropGrowthLogUsecase) *CropGrowthLogHandler {
	return &CropGrowthLogHandler{CropGrowthLogUC: uc}
}

// CreateLog handles the creation of a new crop growth log
func (h *CropGrowthLogHandler) CreateLog(c echo.Context) error {
	log := new(models.CropGrowthLog)
	// Bind request body to log struct
	if err := c.Bind(log); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input: "+err.Error())
	}

	// Validate CropID (ensure it's provided)
	if log.CropID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Crop ID is required")
	}

	// Call usecase to create the log
	err := h.CropGrowthLogUC.CreateLog(log)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create growth log: "+err.Error())
	}

	return c.JSON(http.StatusCreated, log)
}

// GetLogsByCropID handles retrieving all logs for a specific crop
func (h *CropGrowthLogHandler) GetLogsByCropID(c echo.Context) error {
	cropIDStr := c.Param("cropId")
	cropID, err := strconv.ParseUint(cropIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Crop ID format")
	}

	logs, err := h.CropGrowthLogUC.GetLogsByCropID(uint(cropID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve growth logs: "+err.Error())
	}

	return c.JSON(http.StatusOK, logs)
}

// GetLogByID handles retrieving a specific log by its ID
func (h *CropGrowthLogHandler) GetLogByID(c echo.Context) error {
	logIDStr := c.Param("logId")
	logID, err := strconv.ParseUint(logIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Log ID format")
	}

	log, err := h.CropGrowthLogUC.GetLogByID(uint(logID))
	if err != nil {
		// Consider checking for gorm.ErrRecordNotFound for a 404 response
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve growth log: "+err.Error())
	}

	return c.JSON(http.StatusOK, log)
}

// UpdateLog handles updating an existing crop growth log
func (h *CropGrowthLogHandler) UpdateLog(c echo.Context) error {
	logIDStr := c.Param("logId")
	logID, err := strconv.ParseUint(logIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Log ID format")
	}

	logUpdate := new(models.CropGrowthLog)
	if err := c.Bind(logUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input: "+err.Error())
	}

	logUpdate.ID = uint(logID) // Ensure the ID from the path is used

	err = h.CropGrowthLogUC.UpdateLog(logUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update growth log: "+err.Error())
	}

	return c.JSON(http.StatusOK, logUpdate)
}

// DeleteLog handles deleting a crop growth log
func (h *CropGrowthLogHandler) DeleteLog(c echo.Context) error {
	logIDStr := c.Param("logId")
	logID, err := strconv.ParseUint(logIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Log ID format")
	}

	err = h.CropGrowthLogUC.DeleteLog(uint(logID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete growth log: "+err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
