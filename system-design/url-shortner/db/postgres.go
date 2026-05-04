package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Conn *pgx.Conn

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env File")
	}
	connStr := os.Getenv("DATABASE_URL")
	// fmt.Println("DATABASE_URL =", connStr)

	Conn, err = pgx.Connect(context.Background(), connStr)

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Postgres")

}

func CloseDB() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}

func DeleteExpired(ctx context.Context) error {
	_, err := Conn.Exec(ctx, "DELETE FROM urls WHERE expires_at IS NOT NULL AND expires_at < NOW()")
	return err
}
