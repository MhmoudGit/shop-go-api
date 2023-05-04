package auth

import (
	"net/http"
	"time"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var user models.User

// hash password
func hashPassword(password string) (string, error) {
	// Generate a hashed version of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// get user by email
func GetUserByEmail(email string) *gorm.DB {
	user := db.Db.Model(&models.User{}).Where("Email = ?", email).First(&user)
	return user
}

// verify user password
func VerifyPassword(password string) bool {
	//hash password
	return true
}

func GenerateAccessToken(userID int) (string, error) {
	// Define the claims for the token
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate the token using HMAC SHA256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Authenticate the user and retrieve the user ID
	userID := 123 // Replace with the actual user ID

	// Generate an access token for the authenticated user
	accessToken, err := GenerateAccessToken(userID)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	// Return the access token to the client
	w.Header().Set("Authorization", "Bearer "+accessToken)
	w.WriteHeader(http.StatusOK)
}
