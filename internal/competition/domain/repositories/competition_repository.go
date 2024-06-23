package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type CompetitionRepository interface {
	Create(competition *entities.Competition) error
	FindAll() ([]*entities.Competition, error)
	FindById(id uuid.UUID) (*entities.Competition, error)
}
