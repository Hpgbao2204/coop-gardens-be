package models

import (
    "time"
)

type Inventory struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"not null"`  // Tên vật tư
    Category    string    `gorm:"not null"`  // Phân loại (Phân bón, thuốc trừ sâu,...)
    Quantity    float64   `gorm:"not null"`  // Số lượng
    Unit        string    `gorm:"not null"`  // Đơn vị (kg, lít,...)
    Status      string    `gorm:"default:'In Stock'"` // Trạng thái
    CreatedBy   uint      `gorm:"not null"`  // Người nhập kho
    User        User      `gorm:"foreignKey:CreatedBy"`
    LastUpdated time.Time `gorm:"autoUpdateTime"`
}

type CropInventory struct {
	ID          uint      `gorm:"primaryKey"`
	CropID      uint      `gorm:"not null"`
	InventoryID uint      `gorm:"not null"`
	Quantity    float64   `gorm:"not null"` // Lượng vật tư dùng cho cây trồng/mùa vụ
	Crop        Crop      `gorm:"foreignKey:CropID"`
	Inventory   Inventory `gorm:"foreignKey:InventoryID"`
}

type InventoryTransaction struct {
	ID          uint      `gorm:"primaryKey"`
	InventoryID uint      `gorm:"not null"`
	Type        string    `gorm:"not null"` // "import" hoặc "export"
	Quantity    float64   `gorm:"not null"`
	PerformedBy uint      `gorm:"not null"` // Ai nhập/xuất kho
	User        User      `gorm:"foreignKey:PerformedBy"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}