package routes

import (
	"coop-gardens-be/internal/api/handlers"
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"

	"github.com/labstack/echo/v4"
)

func CropRoutes(g *echo.Group, cropHandler *handlers.CropHandler, userRepo *repository.UserRepository) {
	g.Use(middlewares.JWTMiddleware)
    
    // For POST requests, require Admin role
    adminRoute := g.Group("")
    adminRoute.Use(middlewares.RoleMiddleware("Admin", userRepo))
    adminRoute.POST("", cropHandler.CreateCrop)
    
    // GET requests are public (no role restriction)
    g.GET("", cropHandler.GetAllCrops)
}

