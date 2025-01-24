package tests

import (
	"log"

	"github.com/Vergiananta/be-simple-zoom/db/initializers"
	"github.com/Vergiananta/be-simple-zoom/internal/models"

	"github.com/joho/godotenv"
)

// DatabaseRefresh runs fresh migration
func DatabaseRefresh() {
	// Load env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	initializers.ConnectDB()

	// Drop all the tables
	err = initializers.DB.Migrator().DropTable(models.User{}, models.Meeting{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	// Migrate again
	err = initializers.DB.AutoMigrate(models.User{}, models.Meeting{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
