package routes

import (
	"coop-gardens-be/internal/api/handlers"
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"

	"github.com/labstack/echo/v4"
)

func SeasonRoutes(g *echo.Group, seasonHandler *handlers.SeasonHandler, userRepo *repository.UserRepository) {
	// Apply JWT middleware to the entire group
	g.Use(middlewares.JWTMiddleware)

	// GET requests - available to all authenticated users
	g.GET("", seasonHandler.GetAllSeasons)
	g.GET("/:id", seasonHandler.GetSeasonByID)
	g.GET("/:id/crops", seasonHandler.GetSeasonWithCrops)

	// POST, PUT, DELETE - only for Admin users
	adminRoutes := g.Group("")
	adminRoutes.Use(middlewares.RoleMiddleware("Admin", userRepo))
	adminRoutes.POST("", seasonHandler.CreateSeason)
	adminRoutes.PUT("/:id", seasonHandler.UpdateSeason)
	adminRoutes.DELETE("/:id", seasonHandler.DeleteSeason)
}
