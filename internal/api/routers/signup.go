package routers

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func SignupRoutes(g *echo.Group) {
	g.POST("/signup", handlers.Signup)
}
