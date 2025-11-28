package support

import "github.com/hse-trpo-taxi/backend/models"

type SupportUseCase interface {
	GetRecent() ([]*models.SupportModel, error)
}

type SupportRepository interface {
	GetRecent() ([]*models.SupportModel, error)
}
