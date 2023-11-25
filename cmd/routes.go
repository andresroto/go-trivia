package main

import (
	"github.com/andresroto/go-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/facts")

	api.Get("/", handlers.ListFacts)
	api.Get("/:id", handlers.ShowFact)
	api.Post("/", handlers.CreateFact)
	api.Put("/:id", handlers.UpdateFact)
	api.Delete("/:id", handlers.DeleteFact)
}
