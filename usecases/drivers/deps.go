package drivers

import (
	"github.com/hse-trpo-taxi/backend/models"
)

//go:generate mockgen --source=deps.go --destination=mocks/mock.go

type DriverUseCase interface {
	GetDrivers() ([]*models.Driver, error)
	GetDriverById(id uint32) (*models.DriverResponseModel, error)
	CreateDriver(model *models.CreateDriverModel) (*models.Driver, error)
	UpdateDriver(id uint32, model *models.UpdateDriverModel) (*models.Driver, error)
	DeleteDriver(id uint32) error
	GetDriversMeanScore() (*models.DriversMeanScore, error)
	GetDriversCoords() ([]*models.DriverCoords, error)
	GetDriversStat(model *models.DriverStatRequestModel) ([]*models.DriverStat, error)
	GetDriverSchedule(id uint32, req *models.DriverScheduleRequestModel) ([]*models.DriverSchedule, error)
}

type DriverRepository interface {
	GetDrivers() ([]*models.Driver, error)
	GetDriverById(id uint32) (*models.Driver, error)
	CreateDriver(model *models.CreateDriverModel) (*models.Driver, error)
	UpdateDriver(id uint32, model *models.UpdateDriverModel) (*models.Driver, error)
	DeleteDriver(id uint32) error
	GetDriversCoords() ([]*models.DriverCoords, error)
	GetDriversStat(model *models.DriverStatRequestModel) ([]*models.DriverStat, error)
	GetDriverSchedule(id uint32, req *models.DriverScheduleRequestModel) ([]*models.DriverSchedule, error)
}
