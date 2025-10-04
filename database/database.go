// Package database provides database connection management and initialization for the taxi service.
// It handles PostgreSQL database connections, table creation, and provides a global database instance.
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB is the global database connection instance used throughout the application.
// It is initialized by InitDB and should be closed using CloseDB when the application shuts down.
var DB *sql.DB

// InitDB initializes the database connection using the provided data source name.
// It opens a PostgreSQL connection, verifies connectivity with a ping,
// and creates the necessary tables for the taxi service.
// Returns an error if the connection fails or table creation fails.
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("Database connection established")
	return createTables()
}

// createTables creates the necessary database tables for the taxi service.
// It creates three tables: clients, drivers, and cars with their respective schemas.
// The cars table has a foreign key reference to the drivers table.
// Returns an error if any table creation fails.
func createTables() error {
	clientsTable := `
	CREATE TABLE IF NOT EXISTS clients (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		phone VARCHAR(50) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	driversTable := `
	CREATE TABLE IF NOT EXISTS drivers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		phone VARCHAR(50) NOT NULL,
		license_number VARCHAR(50) NOT NULL,
		rating REAL DEFAULT 0.0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	carsTable := `
	CREATE TABLE IF NOT EXISTS cars (
		id SERIAL PRIMARY KEY,
		driver_id INTEGER NOT NULL,
		brand VARCHAR(100) NOT NULL,
		model VARCHAR(100) NOT NULL,
		year INTEGER NOT NULL,
		license_plate VARCHAR(50) NOT NULL,
		color VARCHAR(50) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (driver_id) REFERENCES drivers(id)
	);`

	tables := []string{clientsTable, driversTable, carsTable}
	for _, table := range tables {
		if _, err := DB.Exec(table); err != nil {
			return fmt.Errorf("error creating table: %v", err)
		}
	}

	log.Println("Database tables created successfully")
	return nil
}

// CloseDB safely closes the database connection if it exists.
// This function should be called when the application shuts down
// to ensure proper cleanup of database resources.
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
