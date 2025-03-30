package handlers

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/usecase"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type SeasonHandler struct {
	usecase *usecase.SeasonUsecase
}

func NewSeasonHandler(usecase *usecase.SeasonUsecase) *SeasonHandler {
	return &SeasonHandler{usecase}
}

func (h *SeasonHandler) CreateSeason(c echo.Context) error {
	var req struct {
		Name      string `json:"name" validate:"required"`
		StartDate string `json:"start_date" validate:"required"` // Format: YYYY-MM-DD
		EndDate   string `json:"end_date" validate:"required"`   // Format: YYYY-MM-DD
		Status    string `json:"status"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid start date format. Use YYYY-MM-DD"})
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid end date format. Use YYYY-MM-DD"})
	}

	// Default status if not provided
	status := req.Status
	if status == "" {
		status = "Planning"
	}

	season := &models.Season{
		Name:      req.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    status,
	}

	if err := h.usecase.CreateSeason(season); err != nil {
		log.Printf("Error creating season: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create season: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, season)
}

func (h *SeasonHandler) GetAllSeasons(c echo.Context) error {
	seasons, err := h.usecase.GetAllSeasons()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch seasons"})
	}
	return c.JSON(http.StatusOK, seasons)
}

func (h *SeasonHandler) GetSeasonByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid season ID"})
	}

	season, err := h.usecase.GetSeasonByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch season"})
	}

	return c.JSON(http.StatusOK, season)
}

func (h *SeasonHandler) GetSeasonWithCrops(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid season ID"})
	}

	season, err := h.usecase.GetSeasonWithCrops(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch season"})
	}

	return c.JSON(http.StatusOK, season)
}

func (h *SeasonHandler) UpdateSeason(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid season ID"})
	}

	// First get the existing season
	existingSeason, err := h.usecase.GetSeasonByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Season not found"})
	}

	// Bind request data
	var req struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"` // Format: YYYY-MM-DD
		EndDate   string `json:"end_date"`   // Format: YYYY-MM-DD
		Status    string `json:"status"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Update fields if provided
	if req.Name != "" {
		existingSeason.Name = req.Name
	}

	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid start date format. Use YYYY-MM-DD"})
		}
		existingSeason.StartDate = startDate
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid end date format. Use YYYY-MM-DD"})
		}
		existingSeason.EndDate = endDate
	}

	if req.Status != "" {
		existingSeason.Status = req.Status
	}

	if err := h.usecase.UpdateSeason(existingSeason); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update season"})
	}

	return c.JSON(http.StatusOK, existingSeason)
}

func (h *SeasonHandler) DeleteSeason(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid season ID"})
	}

	if err := h.usecase.DeleteSeason(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete season"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Season deleted successfully"})
}
