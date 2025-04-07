package routes

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

// InventoryRoutes registers inventory-related endpoints.
func InventoryRoutes(g *echo.Group, inventoryHandler *handlers.InventoryHandler) {
	// Route to import inventory
	g.POST("/import", inventoryHandler.ImportInventory)
	// Route to export inventory
	g.POST("/export/:inventory_id", inventoryHandler.ExportInventory)
}
