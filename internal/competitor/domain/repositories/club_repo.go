package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/google/uuid"
)

type ClubRepo interface {
	Create(competitor *entities.Club) error
	FindAll() ([]*entities.Club, error)
	FindById(id uuid.UUID) (*entities.Club, error)
}
