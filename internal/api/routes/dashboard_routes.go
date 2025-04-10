package routes

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func DashboardRoutes(g *echo.Group, dh *handlers.DashboardHandler) {
	// Ví dụ: GET /summary
	g.GET("/summary", dh.GetDashboard)
}
