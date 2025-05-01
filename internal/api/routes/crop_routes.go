package routes

import (
	"coop-gardens-be/internal/api/handlers"
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"

	"github.com/labstack/echo/v4"
)

func CropRoutes(g *echo.Group, cropHandler *handlers.CropHandler, userRepo *repository.UserRepository, growthLogHandler *handlers.CropGrowthLogHandler) {
	g.Use(middlewares.JWTMiddleware)

	// For POST requests, require Admin role
	adminRoute := g.Group("")
	adminRoute.Use(middlewares.RoleMiddleware("Admin", userRepo))
	adminRoute.POST("", cropHandler.CreateCrop)

	// GET requests are public (no role restriction)
	g.GET("", cropHandler.GetAllCrops)
	g.GET(("/season/:season_id/crops"), cropHandler.GetCropsBySeason)
	g.POST("/season/:season_id/crops", cropHandler.AddCropToSeason, middlewares.RoleMiddleware("Admin", userRepo))

	// Growth log routes
	g.POST("/growth-logs", growthLogHandler.CreateLog)
	g.GET("/growth-logs/crop/:cropId", growthLogHandler.GetLogsByCropID)
	g.GET("/growth-logs/:logId", growthLogHandler.GetLogByID)
	g.PUT("/growth-logs/:logId", growthLogHandler.UpdateLog)
	g.DELETE("/growth-logs/:logId", growthLogHandler.DeleteLog)
}
