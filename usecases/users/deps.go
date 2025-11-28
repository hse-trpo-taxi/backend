package users

import "github.com/hse-trpo-taxi/backend/models"

type UserUseCase interface {
	GetCurrent() (*models.UserModel, error)
}
