package models

import "time"

type Driver struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Phone         string    `json:"phone" db:"phone"`
	LicenseNumber string    `json:"license_number" db:"license_number"`
	Rating        float64   `json:"rating" db:"rating"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
