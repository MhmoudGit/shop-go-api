package routers

import (
	"github.com/MhmoudGit/shop-go-api/controllers"
	"github.com/gorilla/mux"
)

func SetupProductRoutes(router *mux.Router) {
	// create a new subrouter for /products
	productsRouter := router.PathPrefix("/products").Subrouter()

	// define the routes for the subrouter
	productsRouter.HandleFunc("", controllers.CreateProduct).Methods("POST")
	productsRouter.HandleFunc("/{CategoryId}", controllers.GetProducts).Methods("GET")
	productsRouter.HandleFunc("/{id}", controllers.GetProduct).Methods("GET")
	productsRouter.HandleFunc("/{id}", controllers.UpdateProduct).Methods("PUT")
	productsRouter.HandleFunc("/{id}", controllers.DeleteProduct).Methods("DELETE")
}
