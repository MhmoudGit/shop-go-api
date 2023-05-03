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

var product models.Product

// get products by category id
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the query string
	CategoryIdStr := mux.Vars(r)["CategoryId"]
	if CategoryIdStr == "" {
		// If the ID parameter is missing, return a 400 Bad Request status code
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// Convert the ID parameter to an integer
	id, err := strconv.Atoi(CategoryIdStr)
	if err != nil {
		// If the ID parameter is not a valid integer, return a 400 Bad Request status code
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// Retrieve the products from the database by categoryID
	var products []models.Product
	err = db.Db.Model(&models.Product{}).Where("category_id = ?", id).Find(&products).Error
	if err != nil {
		// If there was an error fetching categories, return a 500 status code
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	if len(products) == 0 {
		// If no products were found, return a 404 status code
		http.Error(w, "404 No products found", http.StatusNotFound)
		return
	}

	// Return the peoducts json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&products)
}

// get single product by id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the query string
	IdStr := mux.Vars(r)["id"]
	if IdStr == "" {
		// If the ID parameter is missing, return a 400 Bad Request status code
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// Convert the ID parameter to an integers
	id, err := strconv.Atoi(IdStr)
	if err != nil {
		// If the ID parameter is not a valid integer, return a 400 Bad Request status code
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// Retrieve the product from the database by ID
	err = db.Db.Model(&models.Product{}).Where("ID = ?", id).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// If no product was found, return a 404 Not Found status code
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		} else {
			// If there was an error fetching the product, return a 500 Internal Server Error status code
			http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
			return
		}
	}

	// Return product json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&product)
}

// post a product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the product into the database
	err = db.Db.Model(&models.Product{}).Create(&product).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// delete a product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the URL
	idStr := mux.Vars(r)
	id, err := strconv.Atoi(idStr["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Delete the category from the database
	result := db.Db.Model(&models.Product{}).Unscoped().Delete(&product, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Success"))
}

// update a produc
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the query string
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		// If the ID parameter is missing, return a 400 Bad Request status code
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find the product in the database by ID
	var existingProduct models.Product
	result := db.Db.Model(&models.Product{}).First(&existingProduct, idStr)
	if result.Error != nil {
		http.Error(w, "Failed to find product", http.StatusInternalServerError)
		return
	}
	// Update the existing product with the new data
	existingProduct.Name = product.Name
	// Update the product in the database
	result = db.Db.Model(&existingProduct).Updates(product)
	if result.Error != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	// Return the updated product
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&product)
}
