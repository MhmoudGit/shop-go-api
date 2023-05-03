package models

import (
	"gorm.io/gorm"
)

// Database model
type Category struct {
	gorm.Model
	Name     string    `json:"name" gorm:"unique;not null;index"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
