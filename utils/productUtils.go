package utils

import (
	"fmt"
	"net/http"

	"github.com/MhmoudGit/shop-go-api/models"
)

func ParseMultiPartProduct(r *http.Request) models.Product {
	// Parse the request body
	r.ParseMultipartForm(10 << 20)

	name := r.Form.Get("productName")
	price := r.Form.Get("price")
	categoryId := r.Form.Get("categodyId")
	image, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer image.Close()
	UploadImage(image, *handler)

	priceInt := ToInt(price)
	categoryIdInt := ToInt(categoryId)

	imagePath := fmt.Sprintf("/images/%v", handler.Filename)
	product := models.Product{
		Name:       name,
		Price:      priceInt,
		CategoryID: categoryIdInt,
		Image:      imagePath,
	}
	return product
}
