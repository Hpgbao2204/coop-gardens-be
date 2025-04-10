package handlers

import (
	"net/http"
	"strconv"

	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ProductOrderHandler struct {
	usecase *usecase.ProductOrderUsecase
}

func NewProductOrderHandler(usecase *usecase.ProductOrderUsecase) *ProductOrderHandler {
	return &ProductOrderHandler{usecase: usecase}
}

// ---------- Product Handlers ----------
func (h *ProductOrderHandler) CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := h.usecase.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, product)
}

func (h *ProductOrderHandler) GetProducts(c echo.Context) error {
	products, err := h.usecase.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductOrderHandler) GetProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}
	product, err := h.usecase.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

// ---------- Order Handlers ----------
func (h *ProductOrderHandler) CreateOrder(c echo.Context) error {
	var order models.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := h.usecase.CreateOrder(&order); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, order)
}

func (h *ProductOrderHandler) GetOrders(c echo.Context) error {
	orders, err := h.usecase.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *ProductOrderHandler) GetOrderByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}
	order, err := h.usecase.GetOrderByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, order)
}
