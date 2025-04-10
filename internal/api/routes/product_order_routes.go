package routes

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func ProductOrderRoutes(g *echo.Group, ph *handlers.ProductOrderHandler) {
	// Product endpoints
	g.POST("/products", ph.CreateProduct)
	g.GET("/products", ph.GetProducts)
	g.GET("/products/:id", ph.GetProductByID)

	// Order endpoints
	g.POST("/orders", ph.CreateOrder)
	g.GET("/orders", ph.GetOrders)
	g.GET("/orders/:id", ph.GetOrderByID)
}
