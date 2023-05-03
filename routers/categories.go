package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// get all categories
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

// get single category by id
func GetCategoty(w http.ResponseWriter, r *http.Request) {

	// set response header to json
	w.Header().Set("Content-Type", "application/json")

	// Get the ID parameter from the query string
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		// If the ID parameter is missing, return a 400 Bad Request status code
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// Convert the ID parameter to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// If the ID parameter is not a valid integer, return a 400 Bad Request status code
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// Retrieve the category from the database by ID
	var category models.Category
	err = db.Db.Where("ID = ?", id).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// If no category was found, return a 404 Not Found status code
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		} else {
			// If there was an error fetching the category, return a 500 Internal Server Error status code
			http.Error(w, "Failed to fetch category", http.StatusInternalServerError)
			return
		}
	}

	// encode the response to return json, return a 200 status code and the list of categories
	json.NewEncoder(w).Encode(&category)
}

// post a category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the category into the database
	err = db.Db.Create(&category).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
