package main

import (
	"github.com/gofiber/fiber/v2"
	"library-management/book"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetDetailBook)
	app.Post("/api/v1/book", book.AddNewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
