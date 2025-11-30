package handlers

import (
	"log/slog"
	"net/http"

	"github.com/hse-trpo-taxi/backend/usecases/support"
)

type SupportHandler struct {
	supportUS support.SupportUseCase
	lgr       *slog.Logger
}

func NewSupportHandler(supportUS support.SupportUseCase, lgr *slog.Logger) *SupportHandler {
	return &SupportHandler{supportUS, lgr}
}

func (handler *SupportHandler) GetRecent(w http.ResponseWriter, r *http.Request) {
	items, err := handler.supportUS.GetRecent()

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "SupportHandler", err)
		return
	}

	respondWithJSON(w, items, handler.lgr, "SupportHandler")
}
