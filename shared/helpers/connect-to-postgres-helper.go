package helpers

import (
	"context"
	"os"

	"monospec-api/shared/enums"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToPostgres() *pgxpool.Pool {
	databaseUrl := os.Getenv(enums.PostgresDatabaseUrl)

	dbPool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		println("Failed to connect to Postgres")
		println(err)
		panic(err)
	}

  println("Postgres connection is established 123")

	return dbPool
}
