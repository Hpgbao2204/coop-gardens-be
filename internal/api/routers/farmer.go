package routers

import (
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// FarmerRoutes chỉ cho Farmer truy cập
func FarmerRoutes(g *echo.Group, userRepo *repository.UserRepository) {
	g.Use(middlewares.JWTMiddleware)
	g.Use(middlewares.RoleMiddleware("Farmer", userRepo))

	g.GET("/dashboard", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Farmer access granted",
		})
	})

	// Farmer specific endpoints
	g.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Farmer products con cac",
		})
	})
}
