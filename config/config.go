package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the SQLite database connection
func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("development.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	log.Println("Connected to SQLite database")
}
