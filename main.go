package main

import (
	"log"
	"net/http"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/middlewares"
	"github.com/MhmoudGit/shop-go-api/routers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Initialize the database connection
	db.ConnectDB()

	// categories routes:
	router.HandleFunc("/categories", routers.GetCategories).Methods("GET")
	router.HandleFunc("/categories", routers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", routers.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", routers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", routers.DeleteCategoty).Methods("DELETE")
	// products routes:
	router.HandleFunc("/products", routers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{CategoryId}", routers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", routers.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", routers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", routers.DeleteProduct).Methods("DELETE")
	//wrap entire mux with logger middleware
	wrappedMux := middlewares.NewLogger(router)

	log.Printf("server is listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", wrappedMux))

}
