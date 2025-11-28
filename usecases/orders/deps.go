package orders

import "github.com/hse-trpo-taxi/backend/models"

type OrderUseCase interface {
	GetWeekStat() (*models.OrderWeekStat, error)
	GetCurrentStat() (*models.CurrentStat, error)
}
