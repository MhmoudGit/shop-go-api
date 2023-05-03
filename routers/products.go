package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/models"
	"github.com/gorilla/mux"
	// "gorm.io/gorm"
)

// get products by category id
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the query string
	CourseIdStr := mux.Vars(r)["courseId"]
	if CourseIdStr == "" {
		// If the ID parameter is missing, return a 400 Bad Request status code
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// Convert the ID parameter to an integer
	id, err := strconv.Atoi(CourseIdStr)
	if err != nil {
		// If the ID parameter is not a valid integer, return a 400 Bad Request status code
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// Retrieve the category from the database by ID
	var products []models.Product
	db.Db.Where("CategoryID = ?", id).Find(&products)
	if len(products) == 0 {
		// If no products were found, return a 404 status code
		http.Error(w, "404 No products found", http.StatusNotFound)
		return
	}

	// Return the updated category
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&products)
}
