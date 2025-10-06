package handlers

import (
	"encoding/json"
	"github.com/hse-trpo-taxi/backend/errors"
	"github.com/hse-trpo-taxi/backend/models"
	"log/slog"
	"net/http"
)

func respondWithError(w http.ResponseWriter, logger *slog.Logger, statusCode int, handlerName string, err error) {
	if err == nil {
		err = &errors.ValidationError{}
	}

	logger.Error(err.Error(), "handler", handlerName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errMessage := &models.ErrorResponse{}

	switch statusCode {
	case http.StatusBadRequest:
		errMessage.Errors = "bad request"
	case http.StatusUnauthorized:
		errMessage.Errors = "invalid credentials"
	case http.StatusForbidden:
		errMessage.Errors = "forbidden"
	default:
		errMessage.Errors = "internal server error"
	}

	response, _ := json.Marshal(errMessage)
	_, _ = w.Write(response)
}

func respondWithJSON(w http.ResponseWriter, payload interface{}, logger *slog.Logger, handlerName string) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, logger, http.StatusInternalServerError, handlerName, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
