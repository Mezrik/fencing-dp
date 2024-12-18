package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type CompetitionRepository interface {
	Create(competition *entities.Competition) error
	Update(competition *entities.Competition) error
	FindAll() ([]*entities.Competition, error)
	FindById(id uuid.UUID) (*entities.Competition, error)

	FindCategoryById(id uuid.UUID) (*entities.CompetitionCategory, error)
	FindAllCategories() ([]*entities.CompetitionCategory, error)

	FindWeaponById(id uuid.UUID) (*entities.Weapon, error)
	FindAllWeapons() ([]*entities.Weapon, error)

	FindAllGroups(competitionId uuid.UUID) ([]*entities.Group, error)
	FindGroupById(id uuid.UUID) (*entities.Group, error)
}
