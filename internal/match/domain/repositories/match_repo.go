package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/match/domain/entities"
	"github.com/google/uuid"
)

type MatchRepository interface {
	FindAll(groupID uuid.UUID) ([]*entities.Match, error)
	FindById(id uuid.UUID) (*entities.Match, error)
}
