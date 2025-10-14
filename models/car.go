package models

import "time"

// Car represents a vehicle used in the taxi service.
// It contains detailed information about the car and its association with a driver.
// Each car belongs to exactly one driver.
type Car struct {
	// ID is the unique identifier for the car
	ID int `json:"id" db:"id"`
	// DriverID is the foreign key reference to the driver who owns this car
	DriverID int `json:"driver_id" db:"driver_id"`
	// Brand is the manufacturer of the car (e.g., Toyota, Honda)
	Brand string `json:"brand" db:"brand"`
	// Model is the specific model of the car (e.g., Camry, Civic)
	Model string `json:"model" db:"model"`
	// Year is the manufacturing year of the car
	Year int `json:"year" db:"year"`
	// LicensePlate is the unique license plate number of the car
	LicensePlate string `json:"license_plate" db:"license_plate"`
	// Color is the color of the car
	Color string `json:"color" db:"color"`
	// CreatedAt is the timestamp when the car record was created
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt is the timestamp when the car record was last modified
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (model Car) InsertColumns() []string {
	return []string{
		"driver_id",
		"brand",
		"model",
		"year",
		"license_plate",
		"color",
		"created_at",
		"updated_at",
	}
}

type CreateCarModel struct {
	DriverID     int    `json:"driver_id" db:"driver_id"`
	Brand        string `json:"brand" db:"brand"`
	Model        string `json:"model" db:"model"`
	Year         int    `json:"year" db:"year"`
	LicensePlate string `json:"license_plate" db:"license_plate"`
	Color        string `json:"color" db:"color"`
}

type UpdateCarModel struct {
	DriverID int    `json:"driver_id" db:"driver_id"`
	Color    string `json:"color" db:"color"`
}

func (model *CreateCarModel) Validate() bool {
	if model.Year < 1950 {
		return false
	}

	return true
}
