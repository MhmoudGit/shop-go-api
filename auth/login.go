package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/MhmoudGit/shop-go-api/config"
	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var user models.User

// hash password
func HashPassword(password string) (string, error) {
	// Generate a hashed version of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password, err
	}
	return string(hashedPassword), nil
}

// get user by email
func GetUserByEmail(email string) (*models.User, error) {
	result := db.Db.Model(&models.User{}).Where("Email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Return a custom "not found" error
			return nil, fmt.Errorf("wrong email or password")
		}
		return nil, result.Error
	}
	return &user, nil
}

// verify user password
func AuthinticateUser(email, password string) (bool, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Handle error, e.g. return authentication failure
		return false, nil
	}
	// Passwords match
	return true, nil
}

func GenerateAccessToken(userID int, role string) (string, error) {
	// Define the claims for the token
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate the token using HMAC SHA256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Authenticate the user and retrieve the user ID
	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userAuth, err := AuthinticateUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userAuth {
		// Generate an access token for the authenticated user
		accessToken, err := GenerateAccessToken(int(user.ID), user.Role)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
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

		// Return the access token to the client
		w.Header().Set("Authorization", "Bearer "+accessToken)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		// Return the access token to the client
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("wrong email or password"))
	}

}
