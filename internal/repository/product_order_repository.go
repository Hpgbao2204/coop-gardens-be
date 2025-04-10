package repository

import (
	"coop-gardens-be/internal/models"

	"gorm.io/gorm"
)

type ProductOrderRepository struct {
	DB *gorm.DB
}

func NewProductOrderRepository(db *gorm.DB) *ProductOrderRepository {
	return &ProductOrderRepository{DB: db}
}

// ---------- Product methods ----------
func (r *ProductOrderRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductOrderRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductOrderRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// ---------- Order methods ----------
func (r *ProductOrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *ProductOrderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.DB.Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *ProductOrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("Items").Find(&orders).Error
	return orders, err
}
