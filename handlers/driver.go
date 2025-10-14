package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/drivers"
	"log/slog"
	"net/http"
	"strconv"
)

type DriverHandler struct {
	driverUS drivers.DriverUseCase
	lgr      *slog.Logger
}

func NewDriverHandler(driverUS drivers.DriverUseCase, lgr *slog.Logger) *DriverHandler {
	return &DriverHandler{driverUS, lgr}
}

// GetDrivers handles GET /api/drivers requests.
// It retrieves all drivers from the database and returns them as a JSON array.
// Returns HTTP 500 if there's a database error.
func (handler *DriverHandler) GetDrivers(w http.ResponseWriter, r *http.Request) {
	items, err := handler.driverUS.GetDrivers()

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "GetDrivers", err)

		return
	}

	respondWithJSON(w, items, handler.lgr, "GetDrivers")
}

// GetDriverById handles GET /api/drivers/{id} requests.
// It retrieves a specific driver by ID and returns it as JSON.
// Returns HTTP 400 if the ID is invalid, HTTP 404 if the driver is not found,
// or HTTP 500 if there's a database error.
func (handler *DriverHandler) GetDriverById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "GetDriverById", err)

		return
	}

	driver, err := handler.driverUS.GetDriverById(uint32(id))

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "GetDriverById", err)

		return
	}

	respondWithJSON(w, driver, handler.lgr, "GetDriverById")
}

// CreateDriver handles POST /api/drivers requests.
// It creates a new driver with the provided JSON data.
// The created_at and updated_at timestamps are automatically set.
// Returns the created driver with HTTP 201 on success,
// HTTP 400 if the request body is invalid, or HTTP 500 if there's a database error.
func (handler *DriverHandler) CreateDriver(w http.ResponseWriter, r *http.Request) {
	var newDriver *models.CreateDriverModel

	if err := json.NewDecoder(r.Body).Decode(&newDriver); err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "CreateDriver", err)

		return
	}

	driver, err := handler.driverUS.CreateDriver(newDriver)

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "CreateDriver", err)

		return
	}

	respondWithJSON(w, driver, handler.lgr, "CreateDriver")
}

// UpdateDriver handles PUT /api/drivers/{id} requests.
// It updates an existing driver with the provided JSON data.
// The updated_at timestamp is automatically set to the current time.
// Returns the updated driver as JSON on success,
// HTTP 400 if the ID or request body is invalid, or HTTP 500 if there's a database error.
func (handler *DriverHandler) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "UpdateDriver", err)

		return
	}

	var newDriver *models.UpdateDriverModel

	if err := json.NewDecoder(r.Body).Decode(&newDriver); err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "UpdateDriver", err)

		return
	}

	driver, err := handler.driverUS.UpdateDriver(uint32(id), newDriver)

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "UpdateDriver", err)

		return
	}

	respondWithJSON(w, driver, handler.lgr, "UpdateDriver")
}

// DeleteDriver handles DELETE /api/drivers/{id} requests.
// It removes a driver from the database by ID.
// Note: This operation may fail if the driver has associated cars due to foreign key constraints.
// Returns HTTP 204 (No Content) on successful deletion,
// HTTP 400 if the ID is invalid, or HTTP 500 if there's a database error.
func (handler *DriverHandler) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "DeleteDriver", err)

		return
	}

	err = handler.driverUS.DeleteDriver(uint32(id))

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "DeleteDriver", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
