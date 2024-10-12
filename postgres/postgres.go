package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres() *Postgres {
	host := "monospecapistack-rdsnestedstackrdsnest-rds34d05673-b5mbbyvdtfuv.cfiwiiwq0xla.eu-central-1.rds.amazonaws.com"
	port := "5432"
	user := "postgres"
	pswd := "9HxW.CGwtuo^=,mOYSKD^wG2a==oNx"
	dbName := "monospec"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, pswd, dbName)

	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)

		os.Exit(1)
	}

	return &Postgres{Pool: pool}
}
