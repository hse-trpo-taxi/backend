package usecases

import (
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/cars"
)

type CarUseCase struct {
	cars.CarRepository
}

func NewCarUseCase(carRepository cars.CarRepository) *CarUseCase {
	return &CarUseCase{
		CarRepository: carRepository,
	}
}

func (useCase *CarUseCase) GetCars() ([]*models.Car, error) {
	items, err := useCase.CarRepository.GetCars()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (useCase *CarUseCase) GetCarById(id uint32) (*models.Car, error) {
	car, err := useCase.CarRepository.GetCarById(id)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *CarUseCase) CreateCar(model *models.CreateCarModel) (*models.Car, error) {
	car, err := useCase.CarRepository.CreateCar(model)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *CarUseCase) UpdateCar(id uint32, model *models.UpdateCarModel) (*models.Car, error) {
	car, err := useCase.CarRepository.UpdateCar(id, model)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *CarUseCase) DeleteCar(id uint32) error {
	return useCase.CarRepository.DeleteCar(id)
}
