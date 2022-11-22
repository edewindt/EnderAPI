package handlers

import (
	"EnderAPI/env"
	"context"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dburl = env.PgUri()

var ctx = context.Background()

var db, _ = pgxpool.New(ctx, dburl)

type HTTPError struct {
	Status string
	Message string
}

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

// GetCharData godoc
// @Summary      Get a single character's data
// @Description  get character by ID
// @Tags         Characters
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Character ID"
// @Success      200  {object}  Character
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /characters/{id} [get]
func GetCharData(c *fiber.Ctx) error {
	var character []*Character
defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	pgxscan.Select(ctx, db, &character, `SELECT * from characters WHERE id = $1`, c.Params("id"))

	return c.JSON(character[0])
	
	}	
	 
// GetAllCharData godoc
// @Summary      Get all character data
// @Description  get all characters
// @Tags         Characters
// @Accept       json
// @Produce      json
// @Success      200  {object}  Character
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /characters [get]
func GetAllCharData(c *fiber.Ctx) error {
	var characters []*Character

	pgxscan.Select(ctx, db, &characters, `SELECT * from characters ORDER by id`)
		return c.JSON(characters)
}

// GetAllArmyData godoc
// @Summary      Get all army data
// @Description  get all armies
// @Tags         Armies
// @Accept       json
// @Produce      json
// @Success      200  {object}  Army
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /armies [get]
func GetAllArmyData(c * fiber.Ctx) error {
	var armies []*Army

	pgxscan.Select(ctx, db, &armies, `SELECT * from armies ORDER by id`)
		return c.JSON(armies)

}

// GetArmyData godoc
// @Summary      Get a single army's data
// @Description  get army by ID
// @Tags         Armies
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Army ID"
// @Success      200  {object}  Army
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /armies/{id} [get]
func GetArmyData(c * fiber.Ctx) error {
	var army []*Army

	pgxscan.Select(ctx, db, &army, `SELECT * from armies WHERE id = $1`, c.Params("id"))
		return c.JSON(army[0])

}

// GetSpecies godoc
// @Summary      Get a single species data
// @Description  get species by ID
// @Tags         Species
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Species ID"
// @Success      200  {object}  Species
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /species/{id} [get]
func GetSpecies(c *fiber.Ctx) error {
	var species []*Species

	pgxscan.Select(ctx, db, &species, `SELECT * from species WHERE id = $1`, c.Params("id"))
	return c.JSON(species[0])
}

// GetAllSpecies godoc
// @Summary      Get all species data
// @Description  get all species
// @Tags         Species
// @Accept       json
// @Produce      json
// @Success      200  {object}  Species
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /species [get]
func GetAllSpecies(c *fiber.Ctx) error {
	var species []*Species

	pgxscan.Select(ctx, db, &species, `SELECT * from species ORDER by id`)
		return c.JSON(species)
}