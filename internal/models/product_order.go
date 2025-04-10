package models

import "time"

// Product đại diện cho sản phẩm
type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null" json:"stock"`
	FarmerID    string    `gorm:"type:uuid;not null" json:"farmer_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// Order đại diện cho đơn đặt hàng
type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    string      `gorm:"type:uuid;not null" json:"user_id"`
	Total     float64     `json:"total"`
	Status    string      `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

// OrderItem đại diện cho chi tiết đơn hàng (sản phẩm trong đơn)
type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `json:"price"`
}
