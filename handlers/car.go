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

// GetCars handles GET /api/cars requests.
// It retrieves all cars from the database and returns them as a JSON array.
// Returns HTTP 500 if there's a database error.
func GetCars(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, driver_id, brand, model, year, license_plate, color, created_at, updated_at FROM cars")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	cars := []models.Car{}
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.DriverID, &car.Brand, &car.Model, &car.Year, &car.LicensePlate, &car.Color, &car.CreatedAt, &car.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// GetCar handles GET /api/cars/{id} requests.
// It retrieves a specific car by ID and returns it as JSON.
// Returns HTTP 400 if the ID is invalid, HTTP 404 if the car is not found,
// or HTTP 500 if there's a database error.
func GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var car models.Car
	err = database.DB.QueryRow("SELECT id, driver_id, brand, model, year, license_plate, color, created_at, updated_at FROM cars WHERE id = ?", id).
		Scan(&car.ID, &car.DriverID, &car.Brand, &car.Model, &car.Year, &car.LicensePlate, &car.Color, &car.CreatedAt, &car.UpdatedAt)

	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

// CreateCar handles POST /api/cars requests.
// It creates a new car with the provided JSON data.
// The created_at and updated_at timestamps are automatically set.
// The driver_id must reference an existing driver.
// Returns the created car with HTTP 201 on success,
// HTTP 400 if the request body is invalid, or HTTP 500 if there's a database error.
func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car.CreatedAt = time.Now()
	car.UpdatedAt = time.Now()

	result, err := database.DB.Exec("INSERT INTO cars (driver_id, brand, model, year, license_plate, color, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		car.DriverID, car.Brand, car.Model, car.Year, car.LicensePlate, car.Color, car.CreatedAt, car.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	car.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

// UpdateCar handles PUT /api/cars/{id} requests.
// It updates an existing car with the provided JSON data.
// The updated_at timestamp is automatically set to the current time.
// The driver_id must reference an existing driver if changed.
// Returns the updated car as JSON on success,
// HTTP 400 if the ID or request body is invalid, or HTTP 500 if there's a database error.
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car.UpdatedAt = time.Now()

	_, err = database.DB.Exec("UPDATE cars SET driver_id = ?, brand = ?, model = ?, year = ?, license_plate = ?, color = ?, updated_at = ? WHERE id = ?",
		car.DriverID, car.Brand, car.Model, car.Year, car.LicensePlate, car.Color, car.UpdatedAt, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	car.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

// DeleteCar handles DELETE /api/cars/{id} requests.
// It removes a car from the database by ID.
// Returns HTTP 204 (No Content) on successful deletion,
// HTTP 400 if the ID is invalid, or HTTP 500 if there's a database error.
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("DELETE FROM cars WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
