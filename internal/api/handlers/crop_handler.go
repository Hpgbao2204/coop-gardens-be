package handlers

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CropHandler struct {
	usecase *usecase.CropUsecase
}

func NewCropHandler(usecase *usecase.CropUsecase) *CropHandler {
	return &CropHandler{usecase}
}

func (h *CropHandler) CreateCrop(c echo.Context) error {
	var crop models.Crop
	if err := c.Bind(&crop); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := h.usecase.CreateCrop(&crop); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create crop"})
	}
	return c.JSON(http.StatusCreated, crop)
}

func (h *CropHandler) GetAllCrops(c echo.Context) error {
	crops, err := h.usecase.GetAllCrops()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch crops"})
	}
	return c.JSON(http.StatusOK, crops)
}

func (h *CropHandler) GetCropsBySeason(c echo.Context) error {
	seasonID, err := strconv.Atoi(c.Param("season_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid season ID"})
	}

	crops, err := h.usecase.GetCropsBySeason(uint(seasonID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch crops"})
	}

	return c.JSON(http.StatusOK, crops)
}

func (h *CropHandler) AddCropToSeason(c echo.Context) error {
	seasonID, err := strconv.Atoi(c.Param("season_id"))
	if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid season ID"})
	}

	var crop models.Crop
	if err := c.Bind(&crop); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Set the season ID from the URL parameter
	crop.SeasonID = uint(seasonID)

	// Debug log
	log.Printf("Adding crop to season ID: %d, Crop data: %+v", seasonID, crop)

	if err := h.usecase.AddCropToSeason(uint(seasonID), &crop); err != nil {
			log.Printf("Error adding crop to season: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add crop: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, crop)
}