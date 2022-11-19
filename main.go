package main

import (
	"EnderAPI/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/data", handlers.GetAllCharData)
	app.Get("/data/:id", handlers.GetCharData)
	app.Get("/names", handlers.GetNames)
	app.Get("/names/:id", handlers.GetName)
	app.Get("/quotes", handlers.GetQuotes)
	app.Get("/quotes/:id", handlers.GetQuote)
	app.Get("/media", handlers.GetAllMedia)
	app.Get("/media/:id", handlers.GetMedia)
}

func main() {

	app := fiber.New()

		Routes(app)

	app.Listen(":3000")
}