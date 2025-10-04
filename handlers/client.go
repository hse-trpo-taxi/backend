// Package handlers provides HTTP request handlers for the taxi service API.
// It contains handlers for managing clients, drivers, and cars through RESTful endpoints.
// All handlers use JSON for request and response formatting and follow standard HTTP status codes.
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

// GetClients handles GET /api/clients requests.
// It retrieves all clients from the database and returns them as a JSON array.
// Returns HTTP 500 if there's a database error.
func GetClients(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, phone, email, created_at, updated_at FROM clients")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	clients := []models.Client{}
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Phone, &client.Email, &client.CreatedAt, &client.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clients = append(clients, client)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

// GetClient handles GET /api/clients/{id} requests.
// It retrieves a specific client by ID and returns it as JSON.
// Returns HTTP 400 if the ID is invalid, HTTP 404 if the client is not found,
// or HTTP 500 if there's a database error.
func GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	var client models.Client
	err = database.DB.QueryRow("SELECT id, name, phone, email, created_at, updated_at FROM clients WHERE id = ?", id).
		Scan(&client.ID, &client.Name, &client.Phone, &client.Email, &client.CreatedAt, &client.UpdatedAt)

	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

// CreateClient handles POST /api/clients requests.
// It creates a new client with the provided JSON data.
// The created_at and updated_at timestamps are automatically set.
// Returns the created client with HTTP 201 on success,
// HTTP 400 if the request body is invalid, or HTTP 500 if there's a database error.
func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client.CreatedAt = time.Now()
	client.UpdatedAt = time.Now()

	result, err := database.DB.Exec("INSERT INTO clients (name, phone, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		client.Name, client.Phone, client.Email, client.CreatedAt, client.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	client.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

// UpdateClient handles PUT /api/clients/{id} requests.
// It updates an existing client with the provided JSON data.
// The updated_at timestamp is automatically set to the current time.
// Returns the updated client as JSON on success,
// HTTP 400 if the ID or request body is invalid, or HTTP 500 if there's a database error.
func UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client.UpdatedAt = time.Now()

	_, err = database.DB.Exec("UPDATE clients SET name = ?, phone = ?, email = ?, updated_at = ? WHERE id = ?",
		client.Name, client.Phone, client.Email, client.UpdatedAt, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

// DeleteClient handles DELETE /api/clients/{id} requests.
// It removes a client from the database by ID.
// Returns HTTP 204 (No Content) on successful deletion,
// HTTP 400 if the ID is invalid, or HTTP 500 if there's a database error.
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("DELETE FROM clients WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
