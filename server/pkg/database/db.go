package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(DATABASE_URL string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), DATABASE_URL)

	if err != nil {
		log.Fatalf("❌ Failed to connect to a Database: %v", err)
	}

	log.Println("💾 Database connected successfully!")
	return pool
}
