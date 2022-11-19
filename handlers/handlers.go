package handlers

import (
	"EnderAPI/env"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var dburl = env.PgUri()

var dbPool, err = pgxpool.Connect(context.Background(), dburl)

var ID int32
var Name string
var Quote string
var Media string

type Character struct {
	ID int32
	Name string
	Quote string
	Media string
}

var CharData []Character 
func GetCharData(c *fiber.Ctx) error {

	rows, err := dbPool.Query(context.Background(), "select * from characters")
	if err != nil {
		log.Fatal("error while executing query")
	}

	for rows.Next() {
		values, err := rows.Values()

		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		
		ID = values[0].(int32)
		Name = values[1].(string)
		Quote = values[2].(string)
		Media = values[3].(string)

		data := Character {
			ID: ID,
			Name: Name,
			Quote:Quote,
			Media:Media,
		}
		
		CharData = append(CharData, data)
		
	}	
	 
	
	// defer dbPool.Close()	
		fmt.Println(CharData)
	return c.JSON(CharData)


}

func GetAllCharData(c *fiber.Ctx) error {
	return c.SendString("All Character Data")
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