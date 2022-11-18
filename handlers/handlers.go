package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetCharData(c *fiber.Ctx)  {
	 c.SendString("Single Character Data")
}

func GetAllCharData(c *fiber.Ctx) error {
	return c.SendString("All Character Data")
}

func GetNames(c *fiber.Ctx)  {
	c.SendString("All Character Names")
}

func GetName(c *fiber.Ctx)  {
	c.SendString("Single Character Name")
}

func GetQuote(c *fiber.Ctx)  {
	c.SendString("Single Character Quote")
}
func GetQuotes(c *fiber.Ctx)  {
	c.SendString("All Character Quotes")
}
func GetMedia(c *fiber.Ctx)  {
	c.SendString("Single Character Media")
}
func GetAllMedia(c *fiber.Ctx)  {
	c.SendString("All Character media")
}