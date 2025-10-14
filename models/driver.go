package models

import "time"

// Driver represents a taxi driver in the service.
// It contains personal information, licensing details, and performance metrics
// for drivers who provide taxi services.
type Driver struct {
	// ID is the unique identifier for the driver
	ID int `json:"id" db:"id"`
	// Name is the full name of the driver
	Name string `json:"name" db:"name"`
	// Phone is the driver's contact phone number
	Phone string `json:"phone" db:"phone"`
	// LicenseNumber is the driver's license number for verification
	LicenseNumber string `json:"license_number" db:"license_number"`
	// Rating is the driver's average rating from clients (0.0 to 5.0)
	Rating float64 `json:"rating" db:"rating"`
	// CreatedAt is the timestamp when the driver record was created
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt is the timestamp when the driver record was last modified
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateDriverModel struct {
	Name          string  `json:"name" db:"name"`
	Phone         string  `json:"phone" db:"phone"`
	LicenseNumber string  `json:"license_number" db:"license_number"`
	Rating        float64 `json:"rating" db:"rating"`
}

func (model *CreateDriverModel) Validate() bool {
	if model.Rating < 0 {
		return false
	}

	return true
}

type UpdateDriverModel struct {
	Name          string  `json:"name" db:"name"`
	Phone         string  `json:"phone" db:"phone"`
	LicenseNumber string  `json:"license_number" db:"license_number"`
	Rating        float64 `json:"rating" db:"rating"`
}
