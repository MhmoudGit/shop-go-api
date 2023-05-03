package routers

import (
	"github.com/MhmoudGit/shop-go-api/controllers"
	"github.com/gorilla/mux"
)

func SetupCategoriesRoutes(router *mux.Router) {
	// create a new subrouter for /categories
	categoriesRouter := router.PathPrefix("/categories").Subrouter()

	// define the routes for the subrouter
	categoriesRouter.HandleFunc("", controllers.GetCategories).Methods("GET")
	categoriesRouter.HandleFunc("", controllers.CreateCategory).Methods("POST")
	categoriesRouter.HandleFunc("/{id}", controllers.GetCategory).Methods("GET")
	categoriesRouter.HandleFunc("/{id}", controllers.UpdateCategory).Methods("PUT")
	categoriesRouter.HandleFunc("/{id}", controllers.DeleteCategoty).Methods("DELETE")
}
