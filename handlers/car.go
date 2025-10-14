package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/errors"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/cars"
	"log/slog"
	"net/http"
	"strconv"
)

type CarHandler struct {
	carUS cars.CarUseCase
	lgr   *slog.Logger
}

func NewCarHandler(carUS cars.CarUseCase, lgr *slog.Logger) *CarHandler {
	return &CarHandler{carUS, lgr}
}

// GetCars handles GET /api/cars requests.
// It retrieves all cars from the database and returns them as a JSON array.
// Returns HTTP 500 if there's a database error.
func (handler *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	items, err := handler.carUS.GetCars()
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "GetCars", err)

		return
	}

	respondWithJSON(w, items, handler.lgr, "GetCars")
}

// GetCarById handles GET /api/cars/{id} requests.
// It retrieves a specific car by ID and returns it as JSON.
// Returns HTTP 400 if the ID is invalid, HTTP 404 if the car is not found,
// or HTTP 500 if there's a database error.
func (handler *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "GetCarById", err)

		return
	}

	car, err := handler.carUS.GetCarById(uint32(id))

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusNotFound, "GetCarById", err)

		return
	}

	respondWithJSON(w, car, handler.lgr, "GetCarById")
}

// CreateCar handles POST /api/cars requests.
// It creates a new car with the provided JSON data.
// The created_at and updated_at timestamps are automatically set.
// The driver_id must reference an existing driver.
// Returns the created car with HTTP 201 on success,
// HTTP 400 if the request body is invalid, or HTTP 500 if there's a database error.
func (handler *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	var newCar *models.CreateCarModel
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "CreateCar", err)

		return
	}

	if !newCar.Validate() {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "CreateCar", &errors.ValidationError{})

		return
	}

	car, err := handler.carUS.CreateCar(newCar)

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "CreateCar", err)

		return
	}

	respondWithJSON(w, car, handler.lgr, "CreateCar")
}

// UpdateCar handles PUT /api/cars/{id} requests.
// It updates an existing car with the provided JSON data.
// The updated_at timestamp is automatically set to the current time.
// The driver_id must reference an existing driver if changed.
// Returns the updated car as JSON on success,
// HTTP 400 if the ID or request body is invalid, or HTTP 500 if there's a database error.
func (handler *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "UpdateCar", err)

		return
	}

	var updateCar *models.UpdateCarModel
	if err := json.NewDecoder(r.Body).Decode(&updateCar); err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "UpdateCar", err)

		return
	}

	car, err := handler.carUS.UpdateCar(uint32(id), updateCar)

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "UpdateCar", err)

		return
	}

	respondWithJSON(w, car, handler.lgr, "UpdateCar")
}

// DeleteCar handles DELETE /api/cars/{id} requests.
// It removes a car from the database by ID.
// Returns HTTP 204 (No Content) on successful deletion,
// HTTP 400 if the ID is invalid, or HTTP 500 if there's a database error.
func (handler *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "DeleteCar", err)

		return
	}

	err = handler.carUS.DeleteCar(uint32(id))
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "DeleteCar", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
