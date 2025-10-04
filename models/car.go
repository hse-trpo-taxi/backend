package models

import "time"

type Car struct {
	ID             int       `json:"id" db:"id"`
	DriverID       int       `json:"driver_id" db:"driver_id"`
	Brand          string    `json:"brand" db:"brand"`
	Model          string    `json:"model" db:"model"`
	Year           int       `json:"year" db:"year"`
	LicensePlate   string    `json:"license_plate" db:"license_plate"`
	Color          string    `json:"color" db:"color"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
