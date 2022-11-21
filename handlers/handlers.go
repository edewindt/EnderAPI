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

type Army struct {
	ID int32
	Name string
	NotableMembers []string
	Media string
}

type Species struct {
	ID int32
	Name string
	Ramen bool
	Varelse bool
	Behavior string
	Media string
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

func GetAllArmyData(c * fiber.Ctx) error {
	var armies []*Army

	pgxscan.Select(ctx, db, &armies, `SELECT * from armies ORDER by id`)
		return c.JSON(armies)

}

func GetArmyData(c * fiber.Ctx) error {
	var army []*Army

	pgxscan.Select(ctx, db, &army, `SELECT * from armies WHERE id = $1`, c.Params("id"))
		return c.JSON(army[0])

}

func GetSpecies(c *fiber.Ctx) error {
	var species []*Species

	pgxscan.Select(ctx, db, &species, `SELECT * from species WHERE id = $1`, c.Params("id"))
	return c.JSON(species[0])
}

func GetAllSpecies(c *fiber.Ctx) error {
	var species []*Species

	pgxscan.Select(ctx, db, &species, `SELECT * from species ORDER by id`)
		return c.JSON(species)
}