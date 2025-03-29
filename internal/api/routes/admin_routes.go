package routes

import (
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"

	"github.com/labstack/echo/v4"

	"net/http"
)

// AdminRoutes chỉ cho Admin truy cập
func AdminRoutes(g *echo.Group, userRepo *repository.UserRepository) {
	g.Use(middlewares.JWTMiddleware)
	g.Use(middlewares.RoleMiddleware("Admin", userRepo))

	// Add handler for root path
	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Admin root access granted",
		})
	})

	g.GET("/dashboard", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Admin access granted",
		})
	})

	// Add more admin specific routes
	g.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "User list accessible by Admin only",
		})
	})
}
