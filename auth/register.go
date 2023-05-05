package auth

import (
	"encoding/json"
	"fmt"

	// "errors"

	"net/http"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/middlewares"
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
		// Insert the user into the database
		err = db.Db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		// the response
		response := models.UserToResponse(&user)
		// Return a JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := middlewares.GetTokenClaimID(w, r)
	if err != nil {
		fmt.Println(err)
	}
	// Retrieve the user from the database by ID
	err = db.Db.Model(&models.User{}).Where("ID = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// If no user was found, return a 404 Not Found status code
			http.Error(w, "user not found", http.StatusNotFound)
			return
		} else {
			// If there was an error fetching the user, return a 500 Internal Server Error status code
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			return
		}
	}
	// the response
	response := models.UserToResponse(&user)
	// Return the updated category
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Admin"))
}
