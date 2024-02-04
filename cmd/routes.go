package main

import (
	"go-url-short/internal/handlers"

	"github.com/gofiber/fiber"
)

func Routes(app *fiber.App) {
	app.Get("/ping", handlers.Ping)
	app.Get("/:url", handlers.ResolveURL)
	app.Post("/api/v1/", handlers.ShortenURL)
}
