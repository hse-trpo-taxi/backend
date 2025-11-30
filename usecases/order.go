package usecases

import (
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/orders"
)

type OrderUseCase struct {
	orders.OrderRepository
}

func NewOrderUseCase(order orders.OrderRepository) *OrderUseCase {
	return &OrderUseCase{order}
}

func (o OrderUseCase) GetWeekStat() (*models.OrderWeekStat, error) {
	return &models.OrderWeekStat{
		Mon: 13,
		Thu: 11,
		Wed: 12,
		Tue: 14,
		Fri: 20,
		Sat: 25,
		Sun: 43,
	}, nil
}

func (o OrderUseCase) GetCurrentStat() (*models.CurrentStat, error) {
	return &models.CurrentStat{
		Order:     62,
		Free:      30,
		Technical: 8,
	}, nil
}

func (o OrderUseCase) GetLastOrders(driverId int, limit int) ([]*models.OrderDriverList, error) {
	return o.OrderRepository.GetLastOrders(driverId, limit)
}

func (o OrderUseCase) GetCurrentOrder(driverId int) (*models.CurrentOrderInfo, error) {
	return o.OrderRepository.GetCurrentOrder(driverId)
}

func (o OrderUseCase) GetOrderList(skip uint64, limit uint64) ([]*models.OrderClientList, error) {
	return o.OrderRepository.GetOrderList(skip, limit)
}
