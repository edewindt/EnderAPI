package main

import (
	"EnderAPI/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes(app *fiber.App) {
	app.Get("/data", handlers.GetAllCharData)
	app.Get("/data/:id", handlers.GetCharData)
	app.Get("/armies", handlers.GetAllArmyData)
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