package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/orders"
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

// POST /api/drivers/{id}/orders
// body: { "count": <int> }
// returns []OrderDriverList
func (handler *OrderHandler) GetDriverOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "OrderHandler", errors.New("missing driver id"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "OrderHandler", err)
		return
	}

	var req struct {
		Count int `json:"count"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err.Error() != "EOF" {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "OrderHandler", err)
		return
	}
	if req.Count <= 0 {
		req.Count = 10
	}

	ordersList, err := handler.us.GetLastOrders(id, req.Count)
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "OrderHandler", err)
		return
	}
	respondWithJSON(w, ordersList, handler.lgr, "OrderHandler")
}

// GET /api/drivers/{id}/order/current
func (handler *OrderHandler) GetDriverCurrentOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "OrderHandler", errors.New("missing driver id"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "OrderHandler", err)
		return
	}

	ord, err := handler.us.GetCurrentOrder(id)
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "OrderHandler", err)
		return
	}
	if ord == nil {
		// return empty object
		respondWithJSON(w, &models.CurrentOrderInfo{}, handler.lgr, "OrderHandler")
		return
	}
	respondWithJSON(w, ord, handler.lgr, "OrderHandler")
}

// POST /api/orders/list
// body: OrderClientListRequestModel
// returns []OrderClientList
func (handler *OrderHandler) GetOrdersList(w http.ResponseWriter, r *http.Request) {
	var req models.OrderClientListRequestModel
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err.Error() != "EOF" {
		respondWithError(w, handler.lgr, http.StatusBadRequest, "OrderHandler", err)
		return
	}

	// defaults
	if req.Limit == 0 {
		req.Limit = 10
	}

	items, err := handler.us.GetOrderList(req.Skip, req.Limit)
	if err != nil {
		respondWithError(w, handler.lgr, http.StatusInternalServerError, "OrderHandler", err)
		return
	}

	respondWithJSON(w, items, handler.lgr, "OrderHandler")
}
