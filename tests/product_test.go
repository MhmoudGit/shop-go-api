package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MhmoudGit/shop-go-api/routers"
)

func TestCreateProduct(t *testing.T) {
	// Create a new request to the CreateProduct endpoint
	req, err := http.NewRequest("POST", "/api/products", strings.NewReader(`{"productName": "Test Product", "price": 100, "categoryId": 1, "image": "/images/test.png"}`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to capture the response from the endpoint
	rr := httptest.NewRecorder()

	// Call the CreateProduct endpoint with the new request and response recorder
	handler := http.HandlerFunc(routers.CreateProduct)
	handler.ServeHTTP(rr, req)

	// Check that the response status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check that the response body contains the expected JSON data
	expected := `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"productName":"Test Product","price":100,"categoryId":1,"image":"test.png"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
