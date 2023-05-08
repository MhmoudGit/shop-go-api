package main

import (
	"log"
	"net/http"

	"github.com/MhmoudGit/shop-go-api/db"
	"github.com/MhmoudGit/shop-go-api/middlewares"
	"github.com/MhmoudGit/shop-go-api/routers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Initialize the database connection
	db.ConnectDB()

	// auth routers
	routers.SetupAuthRoutes(router)
	// categories routes:
	routers.SetupCategoriesRoutes(router)
	// products routes:
	routers.SetupProductRoutes(router)

	//wrap entire mux with logger middleware
	wrappedMux := middlewares.NewLogger(router)

	log.Printf("server is listening on Port 8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", wrappedMux))

}
