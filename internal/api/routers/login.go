package routers

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

// Sau nay config chinh lai thanh echo.echo
func LoginRoutes(g *echo.Group) {
	g.POST("/login", handlers.Login)
}
