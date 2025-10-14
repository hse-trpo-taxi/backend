package cars

import "github.com/hse-trpo-taxi/backend/models"

type CarUseCase interface {
	GetCars() ([]*models.Car, error)
	GetCarById(id uint32) (*models.Car, error)
	CreateCar(model *models.CreateCarModel) (*models.Car, error)
	UpdateCar(id uint32, model *models.UpdateCarModel) (*models.Car, error)
	DeleteCar(id uint32) error
}

type CarRepository interface {
	GetCars() ([]*models.Car, error)
	GetCarById(id uint32) (*models.Car, error)
	CreateCar(model *models.CreateCarModel) (*models.Car, error)
	UpdateCar(id uint32, model *models.UpdateCarModel) (*models.Car, error)
	DeleteCar(id uint32) error
}
