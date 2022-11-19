package database

import (
	"EnderAPI/env"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DbConnection()  {
	dburl := env.PgUri()
	db, err := pgxpool.Connect(context.Background(), dburl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect %v\n", err)
		os.Exit(1)
	}

return db

}