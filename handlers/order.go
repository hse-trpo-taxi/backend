package handlers

import (
	"github.com/hse-trpo-taxi/backend/usecases/orders"
	"log/slog"
	"net/http"
)

type OrderHandler struct {
	us  orders.OrderUseCase
	lgr *slog.Logger
}

func NewOrderHandler(us orders.OrderUseCase, lgr *slog.Logger) *OrderHandler {
	return &OrderHandler{us: us, lgr: lgr}
}

func (handler *OrderHandler) GetWeekStat(w http.ResponseWriter, r *http.Request) {
	stat, err := handler.us.GetWeekStat()

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "OrderHandler", err)
	}

	respondWithJSON(w, stat, handler.lgr, "OrderHandler")
}

func (handler *OrderHandler) GetCurrentStats(w http.ResponseWriter, r *http.Request) {
	stat, err := handler.us.GetCurrentStat()

	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "OrderHandler", err)
	}

	respondWithJSON(w, stat, handler.lgr, "OrderHandler")
}
