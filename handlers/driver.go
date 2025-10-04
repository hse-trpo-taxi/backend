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
