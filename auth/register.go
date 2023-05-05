package auth

import (
	"encoding/json"
	// "errors"

	"net/http"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/models"
	"gorm.io/gorm"
)

// create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbUser, err := GetUserByEmail(user.Email)
	if dbUser != nil {
		http.Error(w, "try again", http.StatusConflict)
	}
	if err != nil {
		hashedPassword, _ := HashPassword(user.Password)

		user.Password = hashedPassword
		user.Role = "user"
		// Insert the category into the database
		err = db.Db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		// the response
		response := models.GetUser{
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

		// Return a JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
