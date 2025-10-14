package usecases

import (
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/drivers"
)

type DriverUseCase struct {
	drivers.DriverRepository
}

func NewDriverUseCase(driverRepository drivers.DriverRepository) *DriverUseCase {
	return &DriverUseCase{
		DriverRepository: driverRepository,
	}
}

func (useCase *DriverUseCase) GetDrivers() ([]*models.Driver, error) {
	items, err := useCase.DriverRepository.GetDrivers()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (useCase *DriverUseCase) GetDriverById(id uint32) (*models.Driver, error) {
	car, err := useCase.DriverRepository.GetDriverById(id)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *DriverUseCase) CreateDriver(model *models.CreateDriverModel) (*models.Driver, error) {
	car, err := useCase.DriverRepository.CreateDriver(model)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *DriverUseCase) UpdateDriver(id uint32, model *models.UpdateDriverModel) (*models.Driver, error) {
	car, err := useCase.DriverRepository.UpdateDriver(id, model)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *DriverUseCase) DeleteDriver(id uint32) error {
	return useCase.DriverRepository.DeleteDriver(id)
}
