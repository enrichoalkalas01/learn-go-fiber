package schemasql

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID   uint    `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductName string  `gorm:"type:varchar(255);not null" json:"product_name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
	Stock       int     `gorm:"not null" json:"stock"`
	CategoryID  *uint   `gorm:"default:null" json:"category_id"` // Foreign key
}
