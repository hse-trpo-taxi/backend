package handlers

import (
	"github.com/hse-trpo-taxi/backend/usecases/users"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	userUS users.UserUseCase
	lgr    *slog.Logger
}

func NewUserHandler(us users.UserUseCase, lgr *slog.Logger) *UserHandler {
	return &UserHandler{userUS: us, lgr: lgr}
}

func (handler *UserHandler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	usr, err := handler.userUS.GetCurrent()

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "UserHandler", err)
	}

	respondWithJSON(w, usr, handler.lgr, "UserHandler")
}
