package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string `json:"productName" gorm:"unique;not null;index"`
	Price      int    `json:"price" gorm:"not null"`
	CategoryID int    `json:"categoryId"`
	Image      string `json:"image"`
}
