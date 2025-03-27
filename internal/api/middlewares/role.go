package middlewares

import (
	"coop-gardens-be/internal/repository"

	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(requiredRole string, userRepo *repository.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, ok := c.Get("user_id").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found"})
			}

			// Check if user has the required role
			hasRole, err := userRepo.HasRole(userID, requiredRole)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error checking user role"})
			}

			if !hasRole {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied: Missing required role"})
			}

			return next(c)
		}
	}
}
