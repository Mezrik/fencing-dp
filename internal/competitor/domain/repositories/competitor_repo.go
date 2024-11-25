package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/google/uuid"
)

type CompetitorRepo interface {
	Create(competitor *entities.Competitor) error
	BatchCreate(competitors []*entities.Competitor) error
	FindAll() ([]*entities.Competitor, error)
	FindById(id uuid.UUID) (*entities.Competitor, error)
	FindAllByCompetitionId(id uuid.UUID) ([]*entities.Participant, error)
	AssignCompetitor(competitorId uuid.UUID, competitionId uuid.UUID) error
}
