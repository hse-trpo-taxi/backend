package usecases

import "github.com/hse-trpo-taxi/backend/models"

type UserUseCase struct {
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (uc *UserUseCase) GetCurrent() (*models.UserModel, error) {
	return &models.UserModel{
		Name:  "Анастасия",
		Score: 82,
	}, nil
}
