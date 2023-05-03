package routers

import (
	"encoding/json"
	"net/http"

	"example.com/dokan/db"
	"example.com/dokan/models"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

	// set response header to json
	w.Header().Set("Content-Type", "application/json")

	// Retrieve all categories from the database
	var categories []models.Category
	err := db.Db.Find(&categories).Error
	if err != nil {
		// If there was an error fetching categories, return a 500 status code
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	if len(categories) == 0 {
		// If no categories were found, return a 404 status code
		http.Error(w, "404 No categories found", http.StatusNotFound)
		return
	}

	// encode the response to return json, return a 200 status code and the list of categories
	json.NewEncoder(w).Encode(&categories)
}
