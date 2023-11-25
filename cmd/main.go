package main

import (
	"github.com/andresroto/go-trivia/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Db connection
	database.ConnectDb()

	app := fiber.New()

	// Routes
	setupRoutes(app)

	// Run server
	app.Listen(":3000")
}
