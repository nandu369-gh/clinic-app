package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB represents our global database connection pool
var DB *pgxpool.Pool

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "clinic_db"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	// Format connection string for PostgreSQL
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	var err error
	// Establish the connection pool
	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	// Ping the database to ensure connection is live
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}
}
