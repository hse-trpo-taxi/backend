package orders

import "github.com/hse-trpo-taxi/backend/models"

type OrderUseCase interface {
	GetWeekStat() (*models.OrderWeekStat, error)
	GetCurrentStat() (*models.CurrentStat, error)
	GetLastOrders(driverId int, limit int) ([]*models.OrderDriverList, error)
	GetCurrentOrder(driverId int) (*models.CurrentOrderInfo, error)
	GetOrderList() ([]*models.OrderClientList, error)
}

type OrderRepository interface {
	GetLastOrders(driverId int, limit int) ([]*models.OrderDriverList, error)
	GetCurrentOrder(driverId int) (*models.CurrentOrderInfo, error)
	GetOrderList() ([]*models.OrderClientList, error)
}
