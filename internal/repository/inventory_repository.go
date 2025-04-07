package repository

import (
	"coop-gardens-be/internal/models"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db}
}

func (r *InventoryRepository) CreateInventory(inventory *models.Inventory) error {
	return r.db.Create(inventory).Error
}

func (r *InventoryRepository) GetAllInventory() ([]models.Inventory, error) {
	var inventory []models.Inventory
	err := r.db.Find(&inventory).Error
	return inventory, err
}

func (r *InventoryRepository) GetInventoryByID(id uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.First(&inventory, id).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *InventoryRepository) UpdateInventoryQuantity(id uint, quantity float64) error {
	return r.db.Model(&models.Inventory{}).Where("id = ?", id).Update("quantity", quantity).Error
}

func (r *InventoryRepository) CreateTransaction(transaction *models.InventoryTransaction) error {
	return r.db.Create(transaction).Error
}
