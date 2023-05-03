package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"productName" gorm:"unique;not null;index"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`
}
