package routes

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

// CropGrowthLogRoutes registers routes for crop growth log operations
func CropGrowthLogRoutes(group *echo.Group, handler *handlers.CropGrowthLogHandler) {
	// Route to create a new growth log
	group.POST("", handler.CreateLog)

	// Route to get all logs for a specific crop
	group.GET("/crop/:cropId", handler.GetLogsByCropID)

	// Route to get a specific log by its ID
	group.GET("/:logId", handler.GetLogByID)

	// Route to update a specific log by its ID
	group.PUT("/:logId", handler.UpdateLog)

	// Route to delete a specific log by its ID
	group.DELETE("/:logId", handler.DeleteLog)
}
