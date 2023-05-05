package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/MhmoudGit/shop-go-api/config"
	"github.com/golang-jwt/jwt/v5"
)

// decrypt the token
func Decrypt(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return []byte(config.JWTSecret), nil
}

// check token and remove Bearer
func checkToken(token string, w http.ResponseWriter) string {
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials"))
		return ""
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	return token
}

// get the admin token role
func getAdminClaims(token *jwt.Token) (string, error) {
	// Get the role claim from the token
	claims := token.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)
	if !ok {
		// Role claim not found or not a string
		return "", errors.New("invalid role")
	}
	if role != "admin" {
		// Admin does not have the necessary role to access the route
		return "", errors.New("invalid access")
	}
	return role, nil
}

// get the user token role
func getUserClaims(token *jwt.Token) (string, error) {
	// Get the role claim from the token
	claims := token.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)
	if !ok {
		// Role claim not found or not a string
		return "", errors.New("invalid role")
	}
	if role != "user" {
		// User does not have the necessary role to access the route
		return "", errors.New("invalid access")
	}
	return role, nil
}

func GetTokenClaimID(w http.ResponseWriter, r *http.Request) (float64, error) {
	accessToken := r.Header.Get("Authorization")
	accessToken = checkToken(accessToken, w)
	token, err := jwt.Parse(accessToken, Decrypt)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return 0, err
	}
	// Get the id claim from the token
	claims := token.Claims.(jwt.MapClaims)
	id, ok := claims["user_id"].(float64)
	if !ok {
		// id claim not found or not a string
		return 0, errors.New("invalid id")
	}
	return id, nil
}

// middleware to handle all routes based on token and role
func AuthMiddleware(adminRoute http.HandlerFunc, userRoute http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = checkToken(tokenString, w)
		token, err := jwt.Parse(tokenString, Decrypt)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		if token.Valid {
			if adminRoute != nil {
				admin, err := getAdminClaims(token)
				if err != nil && admin != "admin" {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(`{"message": "Invalid access"}`))
					return
				}
				adminRoute.ServeHTTP(w, r)
			}
			if userRoute != nil {
				user, err := getUserClaims(token)
				if err != nil && user != "user" {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(`{"message": "Invalid access"}`))
					return
				}
				userRoute.ServeHTTP(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Authorization Token"))
			return
		}
	})
}
