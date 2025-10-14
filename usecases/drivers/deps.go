package drivers

import "github.com/hse-trpo-taxi/backend/models"

type DriverUseCase interface {
	GetDrivers() ([]*models.Driver, error)
	GetDriverById(id uint32) (*models.Driver, error)
	CreateDriver(model *models.CreateDriverModel) (*models.Driver, error)
	UpdateDriver(id uint32, model *models.UpdateDriverModel) (*models.Driver, error)
	DeleteDriver(id uint32) error
}

type DriverRepository interface {
	GetDrivers() ([]*models.Driver, error)
	GetDriverById(id uint32) (*models.Driver, error)
	CreateDriver(model *models.CreateDriverModel) (*models.Driver, error)
	UpdateDriver(id uint32, model *models.UpdateDriverModel) (*models.Driver, error)
	DeleteDriver(id uint32) error
}
