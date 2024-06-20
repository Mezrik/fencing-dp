package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/core/entities"
	"github.com/google/uuid"
)

type CompetitionRepository interface {
	Create(competition *entities.Competition) (*entities.Competition, error)
	FindAll() ([]*entities.Competition, error)
	FindById(id uuid.UUID) (*entities.Competition, error)
}
