package usecases

import (
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/cars"
	"github.com/hse-trpo-taxi/backend/usecases/drivers"
)

type DriverUseCase struct {
	drivers.DriverRepository
	cars.CarRepository
}

func NewDriverUseCase(driverRepository drivers.DriverRepository, carRepository cars.CarRepository) *DriverUseCase {
	return &DriverUseCase{
		DriverRepository: driverRepository,
		CarRepository:    carRepository,
	}
}

func (useCase *DriverUseCase) GetDrivers() ([]*models.Driver, error) {
	items, err := useCase.DriverRepository.GetDrivers()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (useCase *DriverUseCase) GetDriverById(id uint32) (*models.DriverResponseModel, error) {
	driver, err := useCase.DriverRepository.GetDriverById(id)

	if err != nil {
		return nil, err
	}

	car, err := useCase.CarRepository.GetCarByDriverId(uint32(driver.ID))

	if err != nil {
		return nil, err
	}

	return &models.DriverResponseModel{
		Car: &models.CarDriverModel{
			Number:      car.ID,
			Model:       car.Model,
			Running:     car.Running,
			ScoreSystem: car.ScoreSystem,
			ScoreUsers:  car.ScoreUsers,
		},
		Driver: &models.DriverShortModel{
			Name:        driver.Name,
			Phone:       driver.Phone,
			Status:      driver.Status,
			ScoreDriver: driver.ScoreDriver,
			Passport:    driver.Passport,
			Inn:         driver.Inn,
			Snils:       driver.Snils,
		},
	}, nil
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

func (useCase *DriverUseCase) GetDriversMeanScore() (*models.DriversMeanScore, error) {
	return &models.DriversMeanScore{
		Score:       3.78,
		TimeWaiting: 12,
		TimeFree:    2,
	}, nil
}

func (useCase *DriverUseCase) GetDriversCoords() ([]*models.DriverCoords, error) {
	items, err := useCase.DriverRepository.GetDriversCoords()

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (useCase *DriverUseCase) GetDriversStat(model *models.DriverStatRequestModel) ([]*models.DriverStat, error) {
	items, err := useCase.DriverRepository.GetDriversStat(model)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (useCase *DriverUseCase) GetDriverSchedule(id uint32, req *models.DriverScheduleRequestModel) ([]*models.DriverSchedule, error) {
	schedule, err := useCase.DriverRepository.GetDriverSchedule(id, req)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}
