package main

import (
	"EnderAPI/handlers"
	"fmt"
	"os"

	_ "EnderAPI/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title The Ender API
// @version 1.0
// @description This is an open source REST API for Enders Game
// @contact.name Elias De Windt
// @contact.url https://edewindt.com/contact-me/
// @contact.email elias@edewindt.com

func Routes(app *fiber.App) {

	app.Get("/characters", handlers.GetAllCharData)
	app.Get("/characters/:id", handlers.GetCharData)
	app.Get("/armies", handlers.GetAllArmyData)
	app.Get("/armies/:id", handlers.GetArmyData)
	app.Get("/species", handlers.GetAllSpecies)
	app.Get("/species/:id", handlers.GetSpecies)
	app.Get("/swagger/*", swagger.HandlerDefault)

}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {

	fmt.Println("Started")
	os.Getenv("Hello")
	app := fiber.New()
	app.Static("/", "./public")

	app.Use(recover.New())

// Or extend your config for customization
app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders:  "Origin, Content-Type, Accept",
}))

		Routes(app)

	app.Listen(getPort())
}