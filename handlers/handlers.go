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
	ID int32 `json:"id"`
	Name string `json:"name"`
	Quote string `json:"quote"`
	Media string `json:"media"`
	Audio string `json:"audio"`
	Bio string `json:"bio"`
	Trivia string `json:"trivia"`
	Species string `json:"species"`
}

type Army struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
	NotableMembers []string `json:"notable_members"`
	Media string `json:"media"`
}

type Species struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
	Ramen bool `json:"ramen"`
	Varelse bool `json:"varelse"`
	Behavior string `json:"behavior"`
	Media string `json:"media"`
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