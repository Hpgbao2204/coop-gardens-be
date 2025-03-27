package routers

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group, authHandler *handlers.AuthHandler) {
	g.POST("/auth/signup", authHandler.Signup) 
	g.POST("/auth/login", authHandler.Login)
}
