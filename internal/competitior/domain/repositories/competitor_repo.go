package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/competitior/domain/entities"
	"github.com/google/uuid"
)

type CompetitorRepo interface {
	Create(competitor *entities.Competitor) error
	FindAll() ([]*entities.Competitor, error)
	FindById(id uuid.UUID) (*entities.Competitor, error)
	FindAllByCompetitionId(id uuid.UUID) ([]*entities.Competitor, error)
}
