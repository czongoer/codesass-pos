package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	SKU        string  `gorm:"uniqueIndex;type:varchar(100);not null"`
	Name       string  `gorm:"type:varchar(100);not null"`
	Desc       string  `gorm:"type:varchar(255)"`
	Price      float64 `gorm:"not null"`
	Status     uint    `gorm:"not null"`
	Image      string  `gorm:"type:varchar(255);not null"`
	CategoryID uint    `gorm:"not null"`
	Category   Category
}
