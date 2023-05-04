package routers

import (
	"github.com/MhmoudGit/shop-go-api/auth"
	"github.com/MhmoudGit/shop-go-api/middlewares"
	"github.com/gorilla/mux"
)

func SetupAuthRoutes(router *mux.Router) {
	// create a new subrouter for /login
	authRouter := router.PathPrefix("").Subrouter()

	// define the routes for the subrouter
	authRouter.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	authRouter.HandleFunc("/register", auth.CreateUser).Methods("POST")
	authRouter.HandleFunc("/hello", middlewares.ValidateMiddleware(auth.HelloWorld, nil)).Methods("GET")
}
