package db

import (
	"fmt"

	"github.com/MhmoudGit/shop-go-api/config"
	"github.com/MhmoudGit/shop-go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define a global variable to hold the database connection
var Db *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := config.DbString
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Postgres Sql db is connected")

	// Migrate the database model
	err = Db.AutoMigrate(&models.Category{}, &models.Product{}, &models.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	return Db
}
