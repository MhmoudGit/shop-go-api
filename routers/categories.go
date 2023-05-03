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

var category models.Category

// get all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Retrieve all categories from the database
	var categories []models.Category
	err := db.Db.Model(&models.Category{}).Find(&categories).Error
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

	// Return the updated category
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&categories)
}

// get single category by id
func GetCategory(w http.ResponseWriter, r *http.Request) {
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
	err = db.Db.Model(&models.Category{}).Where("ID = ?", id).First(&category).Error
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

	// Return the updated category
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&category)
}

// post a category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the category into the database
	err = db.Db.Model(&models.Category{}).Create(&category).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// delete a category
func DeleteCategoty(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the URL
	idStr := mux.Vars(r)
	id, err := strconv.Atoi(idStr["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// Delete the category from the database
	result := db.Db.Model(&models.Category{}).Unscoped().Delete(&category, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Success"))
}

// put a category
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the query string
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		// If the ID parameter is missing, return a 400 Bad Request status code
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find the category in the database by ID
	var existingCategory models.Category
	result := db.Db.First(&existingCategory, idStr)
	if result.Error != nil {
		http.Error(w, "Failed to find category", http.StatusInternalServerError)
		return
	}
	// Update the existing category with the new data
	existingCategory.Name = category.Name
	// Update the category in the database
	result = db.Db.Model(&existingCategory).Updates(category)
	if result.Error != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	// Return the updated category
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&category)
}
