package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/database"
	"github.com/hse-trpo-taxi/backend/models"
)

// GetDrivers handles GET /api/drivers requests.
// It retrieves all drivers from the database and returns them as a JSON array.
// Returns HTTP 500 if there's a database error.
func GetDrivers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, phone, license_number, rating, created_at, updated_at FROM drivers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	drivers := []models.Driver{}
	for rows.Next() {
		var driver models.Driver
		if err := rows.Scan(&driver.ID, &driver.Name, &driver.Phone, &driver.LicenseNumber, &driver.Rating, &driver.CreatedAt, &driver.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		drivers = append(drivers, driver)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}

// GetDriver handles GET /api/drivers/{id} requests.
// It retrieves a specific driver by ID and returns it as JSON.
// Returns HTTP 400 if the ID is invalid, HTTP 404 if the driver is not found,
// or HTTP 500 if there's a database error.
func GetDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	var driver models.Driver
	err = database.DB.QueryRow("SELECT id, name, phone, license_number, rating, created_at, updated_at FROM drivers WHERE id = ?", id).
		Scan(&driver.ID, &driver.Name, &driver.Phone, &driver.LicenseNumber, &driver.Rating, &driver.CreatedAt, &driver.UpdatedAt)

	if err != nil {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

// CreateDriver handles POST /api/drivers requests.
// It creates a new driver with the provided JSON data.
// The created_at and updated_at timestamps are automatically set.
// Returns the created driver with HTTP 201 on success,
// HTTP 400 if the request body is invalid, or HTTP 500 if there's a database error.
func CreateDriver(w http.ResponseWriter, r *http.Request) {
	var driver models.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driver.CreatedAt = time.Now()
	driver.UpdatedAt = time.Now()

	result, err := database.DB.Exec("INSERT INTO drivers (name, phone, license_number, rating, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		driver.Name, driver.Phone, driver.LicenseNumber, driver.Rating, driver.CreatedAt, driver.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	driver.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(driver)
}

// UpdateDriver handles PUT /api/drivers/{id} requests.
// It updates an existing driver with the provided JSON data.
// The updated_at timestamp is automatically set to the current time.
// Returns the updated driver as JSON on success,
// HTTP 400 if the ID or request body is invalid, or HTTP 500 if there's a database error.
func UpdateDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	var driver models.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driver.UpdatedAt = time.Now()

	_, err = database.DB.Exec("UPDATE drivers SET name = ?, phone = ?, license_number = ?, rating = ?, updated_at = ? WHERE id = ?",
		driver.Name, driver.Phone, driver.LicenseNumber, driver.Rating, driver.UpdatedAt, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	driver.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

// DeleteDriver handles DELETE /api/drivers/{id} requests.
// It removes a driver from the database by ID.
// Note: This operation may fail if the driver has associated cars due to foreign key constraints.
// Returns HTTP 204 (No Content) on successful deletion,
// HTTP 400 if the ID is invalid, or HTTP 500 if there's a database error.
func DeleteDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("DELETE FROM drivers WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
