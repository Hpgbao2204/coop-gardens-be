package routes

import (
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserRoutes cho người dùng thường
func UserRoutes(g *echo.Group, userRepo *repository.UserRepository) {
	g.Use(middlewares.JWTMiddleware)
	g.Use(middlewares.RoleMiddleware("User", userRepo))

	g.GET("/profile", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "User profile access granted",
		})
	})

	// More user specific routes
	g.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Product catalog for users",
		})
	})
}
