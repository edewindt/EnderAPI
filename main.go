package main

import (
	"EnderAPI/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Get("/audio", handlers.GetAllAudio)
	app.Get("/audio/:id", handlers.GetAudio)

}

func main() {

	app := fiber.New()
	app.Static("/", "./public")

// Or extend your config for customization
app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders:  "Origin, Content-Type, Accept",
}))

		Routes(app)

	app.Listen(":3000")
}