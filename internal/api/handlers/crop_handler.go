package handlers

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/usecase"
	"net/http"

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
