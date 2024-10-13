package repos

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) GetUser() string {
	host := "monospecapistack-rdsnestedstackrdsnest-rds34d05673-b5mbbyvdtfuv.cfiwiiwq0xla.eu-central-1.rds.amazonaws.com"
	port := "5432"
	user := "postgres"
	pswd := "9HxW.CGwtuo^=,mOYSKD^wG2a==oNx"
	dbName := "monospec"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, pswd, dbName)

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	return "nice"
}

func (u *UserRepo) RemoveUser() string {
	return "removed"
}
