package usecases

import (
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/hse-trpo-taxi/backend/usecases/support"
)

type SupportUseCase struct {
	support.SupportRepository
}

func NewSupportUseCase(supportRepository support.SupportRepository) *SupportUseCase {
	return &SupportUseCase{
		supportRepository,
	}
}

func (us *SupportUseCase) GetRecent() ([]*models.SupportModel, error) {
	items, err := us.SupportRepository.GetRecent()

	if err != nil {
		return nil, err
	}

	return items, nil
}
