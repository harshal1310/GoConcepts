package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	databaseURL := "postgres://myuser:password@localhost:5432/mydatabase"

	// Create a connection pool
	dbpool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer dbpool.Close()

	fmt.Println("Connected to PostgreSQL successfully!")

	// Execute a simple query
	var greeting string
	err = dbpool.QueryRow(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&greeting)
	if err != nil {
		log.Fatal("Query error:", err)
	}
	fmt.Println(greeting)
}
