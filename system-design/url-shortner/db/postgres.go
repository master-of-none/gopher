package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env File")
	}
	connStr := os.Getenv("DATABASE_URL")
	// fmt.Println("DATABASE_URL =", connStr)

	Pool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Postgres")

}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
}

func DeleteExpired(ctx context.Context) error {
	_, err := Pool.Exec(ctx, "DELETE FROM urls WHERE expires_at IS NOT NULL AND expires_at < NOW()")
	return err
}
