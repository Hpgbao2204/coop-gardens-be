package models

import (
	"time"
)

type Blog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	AuthorID  string    `gorm:"type:uuid;not null" json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	BlogID    uint      `gorm:"not null" json:"blog_id"`
	Blog      Blog      `gorm:"foreignKey:BlogID" json:"-"`
	AuthorID  string    `gorm:"type:uuid;not null" json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorID" json:"-"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Review struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	InventoryID uint      `gorm:"not null" json:"inventory_id"`
	Inventory   Inventory `gorm:"foreignKey:InventoryID" json:"-"`
	UserID      string    `gorm:"type:uuid;not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
	Rating      int       `gorm:"not null" json:"rating"`
	Comment     string    `json:"comment"`
	CreatedAt   time.Time `json:"created_at"`
}
