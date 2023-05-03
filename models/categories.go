package models

import (
	"time"
)

// Database model
type Category struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

// schema for routers
// type Categories struct {
// 	Categories []Category
// }
