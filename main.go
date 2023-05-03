package main

import (
	"log"
	"net/http"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Initialize the database connection
	db.ConnectDB()

	// categories routes:
	r.HandleFunc("/categories", routers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", routers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", routers.GetCategory).Methods("GET")
	r.HandleFunc("/categories/{id}", routers.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", routers.DeleteCategoty).Methods("DELETE")
	// products routes:
	log.Fatal(http.ListenAndServe(":8000", r))
}
