package database

import (
	"EnderAPI/env"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DbConnection() {
	dburl := env.PgUri()
	dbPool, err := pgxpool.Connect(context.Background(), dburl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect %v\n", err)
		os.Exit(1)
	}
	
	var ID int32
	var Hello string

	rows, err := dbPool.Query(context.Background(), "select * from Hello")
	if err != nil {
		log.Fatal("error while executing query")
	}

	for rows.Next() {
		values, err := rows.Values()

		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		ID = values[0].(int32)
		Hello = values[1].(string)

	}

	fmt.Println(Hello, ID)
	
	defer dbPool.Close()
}