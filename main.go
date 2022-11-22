package main

import (
	"EnderAPI/handlers"

	_ "EnderAPI/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title The Ender API
// @version 1.0
// @description This is an open source REST API for Enders Game
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /

func Routes(app *fiber.App) {
	app.Get("/data", handlers.GetAllCharData)
	app.Get("/data/:id", handlers.GetCharData)
	app.Get("/armies", handlers.GetAllArmyData)
	app.Get("/armies/:id", handlers.GetArmyData)
	app.Get("/species", handlers.GetAllSpecies)
	app.Get("/species/:id", handlers.GetSpecies)
	app.Get("/swagger/*", swagger.HandlerDefault)
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