package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
)

type GroupRepository interface {
	Create(group *entities.Group) error
	// Update(group *entities.Group) error
	// FindAll() ([]*entities.Group, error)
	// FindById(id uuid.UUID) (*entities.Group, error)
}
