package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
	"errors"
)

type ProductOrderUsecase struct {
	repo *repository.ProductOrderRepository
}

func NewProductOrderUsecase(repo *repository.ProductOrderRepository) *ProductOrderUsecase {
	return &ProductOrderUsecase{repo}
}

// ---------- Product usecase ----------
func (u *ProductOrderUsecase) CreateProduct(product *models.Product) error {
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 || product.FarmerID == "" {
		return errors.New("invalid product data")
	}
	return u.repo.CreateProduct(product)
}

func (u *ProductOrderUsecase) GetAllProducts() ([]models.Product, error) {
	return u.repo.GetAllProducts()
}

func (u *ProductOrderUsecase) GetProductByID(id uint) (*models.Product, error) {
	return u.repo.GetProductByID(id)
}

// ---------- Order usecase ----------
func (u *ProductOrderUsecase) CreateOrder(order *models.Order) error {
	if order.UserID == "" {
		return errors.New("user_id is required")
	}

	// Tính toán tổng tiền dựa trên các chi tiết đơn hàng (order items)
	var total float64
	for _, item := range order.Items {
		if item.Quantity <= 0 {
			return errors.New("order item quantity must be positive")
		}
		total += float64(item.Quantity) * item.Price
	}
	order.Total = total

	return u.repo.CreateOrder(order)
}

func (u *ProductOrderUsecase) GetOrderByID(id uint) (*models.Order, error) {
	return u.repo.GetOrderByID(id)
}

func (u *ProductOrderUsecase) GetAllOrders() ([]models.Order, error) {
	return u.repo.GetAllOrders()
}
