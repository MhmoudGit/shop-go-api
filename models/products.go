package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Product_Name string `json:"producName" gorm:"unique;not null;index" validate:"required" required:"true"`
	Price        int    `json:"price" gorm:"not null" validate:"required" required:"true"`
	CategoryID   int    `json:"categoryId" gorm:"column:categoryid;unique" validate:"required" required:"true"`
}
