package usecases

import (
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/clients"
)

type ClientUseCase struct {
	clients.ClientRepository
}

func NewClientUseCase(clientRepository clients.ClientRepository) *ClientUseCase {
	return &ClientUseCase{
		ClientRepository: clientRepository,
	}
}

func (useCase *ClientUseCase) GetClients() ([]*models.Client, error) {
	items, err := useCase.ClientRepository.GetClients()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (useCase *ClientUseCase) GetClientById(id uint32) (*models.Client, error) {
	car, err := useCase.ClientRepository.GetClientById(id)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *ClientUseCase) CreateClient(model *models.CreateClientModel) (*models.Client, error) {
	car, err := useCase.ClientRepository.CreateClient(model)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *ClientUseCase) UpdateClient(id uint32, model *models.UpdateClientModel) (*models.Client, error) {
	car, err := useCase.ClientRepository.UpdateClient(id, model)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (useCase *ClientUseCase) DeleteCar(id uint32) error {
	return useCase.ClientRepository.DeleteClient(id)
}
