package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

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

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
