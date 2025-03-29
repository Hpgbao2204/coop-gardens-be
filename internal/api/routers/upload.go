package routers

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

// UploadRoutes khai báo route cho upload ảnh
func UploadRoutes(g *echo.Group, uploadHandler *handlers.UploadHandler) {
	g.POST("/image", uploadHandler.UploadImage)
}
