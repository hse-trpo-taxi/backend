package clients

import "github.com/hse-trpo-taxi/backend/models"

type ClientUseCase interface {
	GetClients() ([]*models.Client, error)
	GetClientById(id uint32) (*models.Client, error)
	CreateClient(model *models.CreateClientModel) (*models.Client, error)
	UpdateClient(id uint32, model *models.UpdateClientModel) (*models.Client, error)
	DeleteClient(id uint32) error
}

type ClientRepository interface {
	GetClients() ([]*models.Client, error)
	GetClientById(id uint32) (*models.Client, error)
	CreateClient(model *models.CreateClientModel) (*models.Client, error)
	UpdateClient(id uint32, model *models.UpdateClientModel) (*models.Client, error)
	DeleteClient(id uint32) error
}
