// Package handlers provides HTTP request handlers for the taxi service API.
// It contains handlers for managing clients, drivers, and cars through RESTful endpoints.
// All handlers use JSON for request and response formatting and follow standard HTTP status codes.
package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/clients"
	"log/slog"
	"net/http"
	"strconv"
)

type ClientHandler struct {
	clientUS clients.ClientUseCase
	lgr      *slog.Logger
}

func NewClientHandler(clientUS clients.ClientUseCase, lgr *slog.Logger) *ClientHandler {
	return &ClientHandler{clientUS, lgr}
}

// GetClients handles GET /api/clients requests.
// It retrieves all clients from the database and returns them as a JSON array.
// Returns HTTP 500 if there's a database error.
func (handler *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	items, err := handler.clientUS.GetClients()

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "GetClients", err)

		return
	}

	respondWithJSON(w, items, handler.lgr, "GetClients")
}

// GetClientById handles GET /api/clients/{id} requests.
// It retrieves a specific client by ID and returns it as JSON.
// Returns HTTP 400 if the ID is invalid, HTTP 404 if the client is not found,
// or HTTP 500 if there's a database error.
func (handler *ClientHandler) GetClientById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "GetClientById", err)

		return
	}

	client, err := handler.clientUS.GetClientById(uint32(id))

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusNotFound, "GetClientById", err)

		return
	}

	respondWithJSON(w, client, handler.lgr, "GetClientById")
}

// CreateClient handles POST /api/clients requests.
// It creates a new client with the provided JSON data.
// The created_at and updated_at timestamps are automatically set.
// Returns the created client with HTTP 201 on success,
// HTTP 400 if the request body is invalid, or HTTP 500 if there's a database error.
func (handler *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var newClient *models.CreateClientModel
	if err := json.NewDecoder(r.Body).Decode(&newClient); err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "CreateClient", err)

		return
	}

	client, err := handler.clientUS.CreateClient(newClient)

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "CreateClient", err)

		return
	}

	respondWithJSON(w, client, handler.lgr, "CreateClient")
}

// UpdateClient handles PUT /api/clients/{id} requests.
// It updates an existing client with the provided JSON data.
// The updated_at timestamp is automatically set to the current time.
// Returns the updated client as JSON on success,
// HTTP 400 if the ID or request body is invalid, or HTTP 500 if there's a database error.
func (handler *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "UpdateClient", err)

		return
	}

	var newClient *models.UpdateClientModel
	if err := json.NewDecoder(r.Body).Decode(&newClient); err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "UpdateClient", err)

		return
	}

	client, err := handler.clientUS.UpdateClient(uint32(id), newClient)

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "UpdateClient", err)

		return
	}

	respondWithJSON(w, client, handler.lgr, "UpdateClient")
}

// DeleteClient handles DELETE /api/clients/{id} requests.
// It removes a client from the database by ID.
// Returns HTTP 204 (No Content) on successful deletion,
// HTTP 400 if the ID is invalid, or HTTP 500 if there's a database error.
func (handler *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "Delete client", err)

		return
	}

	err = handler.clientUS.DeleteClient(uint32(id))
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "Delete client", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
