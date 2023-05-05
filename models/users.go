package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null;index"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null;index"`
}

type GetUser struct {
	gorm.Model
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

// user get schema
func UserToResponse(user *User) GetUser {
	return GetUser{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}
}
