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

	r.HandleFunc("/categories", routers.GetCategories)
	r.HandleFunc("/category/{id}", routers.GetCategoty)
	log.Fatal(http.ListenAndServe(":8000", r))
}
