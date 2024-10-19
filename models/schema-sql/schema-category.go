package schemasql

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey" json:"id"`
	CategoryName string    `gorm:"type:varchar(255);not null" json:"category_name"`
	Products     []Product `gorm:"foreignKey:CategoryID" json:"products"`
}
