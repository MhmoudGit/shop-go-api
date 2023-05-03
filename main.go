package main

import (
	"log"
	"net/http"

	"example.com/dokan/db"
	"example.com/dokan/routers"
)

func main() {
	// Initialize the database connection
	db.ConnectDB()

	http.HandleFunc("/categories", routers.GetCategories)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
