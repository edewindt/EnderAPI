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
	Audio string
	Bio string
	Trivia string
	Species string
}


func GetCharData(c *fiber.Ctx) error {
	var character []*Character

	pgxscan.Select(ctx, db, &character, `SELECT * from characters WHERE id = $1`, c.Params("id"))
		return c.JSON(character[0])

	}	
	 

func GetAllCharData(c *fiber.Ctx) error {
	var characters []*Character

	pgxscan.Select(ctx, db, &characters, `SELECT * from characters ORDER by id`)
		return c.JSON(characters)
}
