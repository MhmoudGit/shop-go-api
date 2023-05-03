package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Product_Name string `json:"name" gorm:"unique;not null;index"`
	Price        int    `json:"price" gorm:"not null"`
	CategoryID   int    `json:"categoryId" gorm:"column:categoryid"`
}
