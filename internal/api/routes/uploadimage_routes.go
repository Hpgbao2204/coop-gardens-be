package routes

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func UploadImageRoutes(g *echo.Group, uploadHandler *handlers.UploadImageHandler) {
	g.POST("/image", uploadHandler.UploadImage)
}
