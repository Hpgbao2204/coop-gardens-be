package handlers

import (
    "net/http"
    "strconv"

    "coop-gardens-be/internal/models"
    "coop-gardens-be/internal/usecase"

    "github.com/labstack/echo/v4"
)

type InventoryHandler struct {
    Usecase *usecase.InventoryUsecase
}

func NewInventoryHandler(usecase *usecase.InventoryUsecase) *InventoryHandler {
    return &InventoryHandler{Usecase: usecase}
}

// ImportInventory handles inventory import requests.
func (h *InventoryHandler) ImportInventory(c echo.Context) error {
    var inventory models.Inventory
    if err := c.Bind(&inventory); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    if err := h.Usecase.ImportInventory(&inventory); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, inventory)
}

type exportRequest struct {
    Quantity    float64 `json:"quantity"`
    PerformedBy uint    `json:"performed_by"`
}

// ExportInventory handles inventory export requests.
func (h *InventoryHandler) ExportInventory(c echo.Context) error {
    idParam := c.Param("inventory_id")
    inventoryID, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid inventory ID"})
    }

    var req exportRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    if err := h.Usecase.ExportInventory(uint(inventoryID), req.Quantity, req.PerformedBy); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Export inventory success"})
}