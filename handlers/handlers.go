package handlers

import (
	"EnderAPI/env"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dburl = env.PgUri()

var ctx = context.Background()

var db, _ = pgxpool.New(ctx, dburl)

type Character struct {
	ID int32
	Name string
	Quote string
	Media string
}

func GetCharData(c *fiber.Ctx) error {
	var character []*Character

	pgxscan.Select(ctx, db, &character, `SELECT * from characters WHERE id = $1`, c.Params("id"))
		return c.JSON(character)

	}	
	 

func GetAllCharData(c *fiber.Ctx) error {
	var characters []*Character

	pgxscan.Select(ctx, db, &characters, `SELECT * from characters`)
		return c.JSON(characters)
}

func GetNames(c *fiber.Ctx) error {
	return c.SendString("All Character Names")
}

func GetName(c *fiber.Ctx) error {
	return c.SendString("Single Character Name")
}

func GetQuote(c *fiber.Ctx) error {
	return c.SendString("Single Character Quote")
}
func GetQuotes(c *fiber.Ctx) error {
	return c.SendString("All Character Quotes")
}
func GetMedia(c *fiber.Ctx) error {
	return c.SendString("Single Character Media")
}
func GetAllMedia(c *fiber.Ctx) error {
	return c.SendString("All Character media")
}