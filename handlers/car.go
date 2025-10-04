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
