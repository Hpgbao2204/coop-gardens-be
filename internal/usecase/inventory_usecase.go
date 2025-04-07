package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
	"fmt"
)

type InventoryUsecase struct {
	repo *repository.InventoryRepository
}

func NewInventoryUsecase(repo *repository.InventoryRepository) *InventoryUsecase {
	return &InventoryUsecase{repo}
}

// 📌 Nhập kho (Import)
func (uc *InventoryUsecase) ImportInventory(inventory *models.Inventory) error {
	return uc.repo.CreateInventory(inventory)
}

// 📌 Xuất kho (Export)
func (uc *InventoryUsecase) ExportInventory(inventoryID uint, quantity float64, performedBy uint) error {
	inventory, err := uc.repo.GetInventoryByID(inventoryID)
	if err != nil {
		return err
	}

	if inventory.Quantity < quantity {
		return fmt.Errorf("not enough stock")
	}

	// Cập nhật số lượng tồn kho
	newQuantity := inventory.Quantity - quantity
	err = uc.repo.UpdateInventoryQuantity(inventoryID, newQuantity)
	if err != nil {
		return err
	}

	// Lưu lịch sử giao dịch
	transaction := &models.InventoryTransaction{
		InventoryID: inventoryID,
		Type:        "export",
		Quantity:    quantity,
		PerformedBy: performedBy,
	}
	return uc.repo.CreateTransaction(transaction)
}
