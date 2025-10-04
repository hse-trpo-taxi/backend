package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phone TEXT NOT NULL,
		email TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	driversTable := `
	CREATE TABLE IF NOT EXISTS drivers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phone TEXT NOT NULL,
		license_number TEXT NOT NULL,
		rating REAL DEFAULT 0.0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	carsTable := `
	CREATE TABLE IF NOT EXISTS cars (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		driver_id INTEGER NOT NULL,
		brand TEXT NOT NULL,
		model TEXT NOT NULL,
		year INTEGER NOT NULL,
		license_plate TEXT NOT NULL,
		color TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
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
