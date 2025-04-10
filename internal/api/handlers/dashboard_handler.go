package handlers

import (
	"net/http"

	"coop-gardens-be/internal/usecase"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	usecase *usecase.DashboardUsecase
}

func NewDashboardHandler(usecase *usecase.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{usecase: usecase}
}

func (h *DashboardHandler) GetDashboard(c echo.Context) error {
	summary, err := h.usecase.GetDashboardSummary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, summary)
}
