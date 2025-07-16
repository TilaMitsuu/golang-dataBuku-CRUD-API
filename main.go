package main

import (
	"book-api/config"
	"book-api/models"
	"book-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Koneksi DB
	config.ConnectDB()

	// Auto migrate tabel
	config.DB.AutoMigrate(&models.User{}, &models.Book{})

	// Setup routes
	routes.Setup(app)

	app.Listen(":3000")
}