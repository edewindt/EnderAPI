package main

import (
	"EnderAPI/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/data", handlers.GetAllCharData)
	app.Get("/data/:id", handlers.GetAllCharData)
	app.Get("/names", handlers.GetAllCharData)
	app.Get("/names/:id", handlers.GetAllCharData)
	app.Get("/quotes", handlers.GetAllCharData)
	app.Get("/quotes/:id", handlers.GetAllCharData)
	app.Get("/media", handlers.GetAllCharData)
	app.Get("/media/:id", handlers.GetAllCharData)
}

func main() {

	app := fiber.New()

		Routes(app)

	app.Listen(":3000")
}