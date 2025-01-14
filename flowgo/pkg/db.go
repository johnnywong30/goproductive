package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// Connect initializes a connection to the PostgreSQL database and returns the *sql.DB.
func Connect() (*sql.DB, error) {
	// Replace with your PostgreSQL connection string
	// connStr := "postgres://jwong:121ilikeCake!@localhost:5432/jwong?sslmode=disable"
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
        "dev",
        "dev",
        "localhost",
        5432,
        "jwong")

	// Open the connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Successfully connected to the database!")

	// Optional: Double-check with a query (e.g., fetching version)
	var version string
	err = db.QueryRow("SELECT version();").Scan(&version)
	if err != nil {
		log.Fatal("Error fetching version: ", err)
	}
	log.Println("PostgreSQL version:", version)
	return db, nil
}

// // CreateTable initializes the required table if not already present.
// func CreateTable(db *sql.DB) error {
// 	// Create a simple "users" table if not exists
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		name VARCHAR(100) NOT NULL,
// 		email VARCHAR(100) NOT NULL UNIQUE
// 	);`

// 	_, err := db.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("failed to create users table: %w", err)
// 	}

// 	log.Println("Table created or already exists.")
// 	return nil
// }
