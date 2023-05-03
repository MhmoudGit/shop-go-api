package main

import (
	"log"
	"net/http"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/routers"
	"github.com/MhmoudGit/shop-go-api/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Initialize the database connection
	db.ConnectDB()

	// test upload file
	r.HandleFunc("/upload", utils.UploadFile).Methods("POST")
	// categories routes:
	r.HandleFunc("/categories", routers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", routers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", routers.GetCategory).Methods("GET")
	r.HandleFunc("/categories/{id}", routers.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", routers.DeleteCategoty).Methods("DELETE")
	// products routes:
	r.HandleFunc("/products", routers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{CategoryId}", routers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", routers.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", routers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", routers.DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
