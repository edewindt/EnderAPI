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
}

type Name struct {
	Name string
}

type Quote struct {
	Quote string
}

type Media struct {
	Media string
}

type Audio struct {
	Audio string
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

func GetNames(c *fiber.Ctx) error {
	var names []*Name

	pgxscan.Select(ctx,db, &names, `SELECT name from characters`)
	return c.JSON(names)
}

func GetName(c *fiber.Ctx) error {
	var name []*Name

	pgxscan.Select(ctx, db, &name, `SELECT name from characters WHERE id = $1`, c.Params("id"))
	return c.JSON(name[0])
}

func GetQuote(c *fiber.Ctx) error {
	var quote []*Quote

	pgxscan.Select(ctx, db, &quote, `SELECT quote from characters WHERE id = $1`, c.Params("id"))
	return c.JSON(quote[0])
}
func GetQuotes(c *fiber.Ctx) error {
	var quotes []*Quote

	pgxscan.Select(ctx, db, &quotes, `SELECT quote from characters`)
	return c.JSON(quotes)
}
func GetMedia(c *fiber.Ctx) error {
	var media []*Media

	pgxscan.Select(ctx, db, &media, `SELECT media from characters WHERE id = $1`, c.Params("id"))
	return c.JSON(media)
}
func GetAllMedia(c *fiber.Ctx) error {
	var medias []*Media

	pgxscan.Select(ctx, db, &medias, `SELECT media from characters`)
	return c.JSON(medias)
}

func GetAudio(c *fiber.Ctx) error {
	var audio []*Audio

	pgxscan.Select(ctx, db, &audio, `SELECT audio from characters WHERE id = $1`, c.Params("id"))
	return c.JSON(audio)
}

func GetAllAudio(c *fiber.Ctx) error {
	var audios []*Audio

	pgxscan.Select(ctx, db, &audios, `SELECT audio from characters`)
	return c.JSON(audios)
}

