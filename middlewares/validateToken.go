package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MhmoudGit/shop-go-api/config"
	"github.com/golang-jwt/jwt/v5"
)

func decrypt(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return []byte(config.JWTSecret), nil
}

func checkToken(token string, w http.ResponseWriter) string {
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials"))
		return ""
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	return token
}

func getAdminClaims(token *jwt.Token, w http.ResponseWriter) bool {
	// Get the role claim from the token
	claims := token.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)
	if !ok {
		// Role claim not found or not a string
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	if role != "admin" {
		// User does not have the necessary role to access the route
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "Invalid access"}`))
		return false
	}
	return true
}

func getUserClaims(token *jwt.Token, w http.ResponseWriter) bool {
	// Get the role claim from the token
	claims := token.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)
	if !ok {
		// Role claim not found or not a string
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	if role != "user" {
		// User does not have the necessary role to access the route
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "Invalid access"}`))
		return false
	}
	return true
}

func ValidateMiddleware(adminRoute http.HandlerFunc, userRoute http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = checkToken(tokenString, w)
		token, err := jwt.Parse(tokenString, decrypt)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		if token.Valid {
			admin := getAdminClaims(token, w)
			if admin {
				adminRoute.ServeHTTP(w, r)
			}
			user := getUserClaims(token, w)
			if user {
				userRoute.ServeHTTP(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Authorization Token"))
			return
		}
	})
}
