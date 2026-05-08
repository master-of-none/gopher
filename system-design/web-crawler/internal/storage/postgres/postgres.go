package postgres

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
		panic("Error Loading .env File")
	}

	connStr := os.Getenv("DATABASE_URL")

	Pool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		panic(err)
	}

	if err := Pool.Ping(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostGres")
	RunMigrations()
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
}

func RunMigrations() {
	_, err := Pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS urls (
			id BIGSERIAL PRIMARY KEY,
			url TEXT UNIQUE NOT NULL,
			status TEXT DEFAULT 'pending',
			depth INT DEFAULT 0,
			discovered_at TIMESTAMP DEFAULT NOW(),
			crawled_at TIMESTAMP
		)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migration ran")
}
