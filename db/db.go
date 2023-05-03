package db

import (
	"fmt"

	"example.com/dokan/config"
	"example.com/dokan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define a global variable to hold the database connection
var Db *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Postgres Sql db is connected")

	// Migrate the database model
	err = Db.AutoMigrate(&models.Category{})
	if err != nil {
		panic("failed to migrate database")
	}

	return Db
}
