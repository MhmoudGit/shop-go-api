package routers

import (
	"github.com/MhmoudGit/shop-go-api/controllers"
	"github.com/MhmoudGit/shop-go-api/middlewares"
	"github.com/gorilla/mux"
)

func SetupCategoriesRoutes(router *mux.Router) {
	// create a new subrouter for /categories
	categoriesRouter := router.PathPrefix("/categories").Subrouter()

	// define the routes for the subrouter
	categoriesRouter.HandleFunc("", controllers.GetCategories).Methods("GET")
	categoriesRouter.HandleFunc("", middlewares.AuthMiddleware(controllers.CreateCategory, nil)).Methods("POST")
	categoriesRouter.HandleFunc("/{id}", controllers.GetCategory).Methods("GET")
	categoriesRouter.HandleFunc("/{id}", middlewares.AuthMiddleware(controllers.UpdateCategory, nil)).Methods("PUT")
	categoriesRouter.HandleFunc("/{id}", middlewares.AuthMiddleware(controllers.DeleteCategoty, nil)).Methods("DELETE")
}
