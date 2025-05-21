package routes

import (
	"coop-gardens-be/internal/api/handlers"
	"coop-gardens-be/internal/api/middlewares"
	"coop-gardens-be/internal/repository"

	"github.com/labstack/echo/v4"
)

// CommonRoutes định nghĩa các routes không phụ thuộc vào role
func CommonRoutes(g *echo.Group, userRepo *repository.UserRepository) {
	// Chỉ yêu cầu xác thực JWT, không yêu cầu role cụ thể
	g.Use(middlewares.JWTMiddleware)

	userHandler := handlers.NewUserHandler(userRepo)

	// Endpoint profile chung cho mọi role
	g.GET("/profile", userHandler.GetUserProfile)
}
