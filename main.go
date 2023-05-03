package main

import (
	"log"
	"net/http"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/routers"
)

func main() {
	// Initialize the database connection
	db.ConnectDB()

	http.HandleFunc("/categories", routers.GetCategories)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
