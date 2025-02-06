package main

import (
	"log"
	"github.com/adnanibrahi0102/seo-analyzer-golang/config"
	"github.com/adnanibrahi0102/seo-analyzer-golang/models"
	"github.com/adnanibrahi0102/seo-analyzer-golang/routes"
)

func main() {
	// load environment variables and connect to the database & redis
	config.LoadENV()
	config.ConnectToDB()
	config.ConnectToRedis()

	err := config.DB.AutoMigrate(&models.SEO_REPORT{})

	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}

	router := routes.SetupRouter()
	log.Println("🚀 Server running on http://localhost:8080")
	router.Run(":8080")

}
