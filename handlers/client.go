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
