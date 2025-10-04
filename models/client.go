// Package models defines the data structures used throughout the taxi service.
// These models represent the core entities: clients, drivers, and cars.
package models

import "time"

// Client represents a taxi service customer.
// It contains personal information and contact details for a client
// who can request taxi services.
type Client struct {
	// ID is the unique identifier for the client
	ID int `json:"id" db:"id"`
	// Name is the full name of the client
	Name string `json:"name" db:"name"`
	// Phone is the client's contact phone number
	Phone string `json:"phone" db:"phone"`
	// Email is the client's email address
	Email string `json:"email" db:"email"`
	// CreatedAt is the timestamp when the client record was created
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt is the timestamp when the client record was last modified
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
