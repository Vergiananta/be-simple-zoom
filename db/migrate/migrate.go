package main

import (
	"log"

	"github.com/Vergiananta/be-simple-zoom/config"
	"github.com/Vergiananta/be-simple-zoom/db/initializers"
	"github.com/Vergiananta/be-simple-zoom/internal/models"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.Migrator().DropTable(models.User{}, models.Meeting{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	err = initializers.DB.AutoMigrate(models.User{}, models.Meeting{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
