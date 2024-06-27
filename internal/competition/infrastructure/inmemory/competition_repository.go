package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InMemoryCompetitionRepository struct {
	db *gorm.DB
}

func NewInMemoryCompetitionRepository(db *gorm.DB) repositories.CompetitionRepository {
	return &InMemoryCompetitionRepository{db: db}
}

func (repo InMemoryCompetitionRepository) Create(competition *entities.Competition) error {
	dbCompetition := repo.marshalCompetition(competition)

	return repo.db.Create(&dbCompetition).Error
}

func (repo InMemoryCompetitionRepository) FindById(id uuid.UUID) (*entities.Competition, error) {
	var dbCompetition models.CompetitionModel

	if err := repo.db.Preload("Category").Preload("Weapon").First(&dbCompetition, id).Error; err != nil {
		return nil, err
	}

	return repo.unmarshalCompetition(&dbCompetition), nil
}

func (repo InMemoryCompetitionRepository) FindAll() ([]*entities.Competition, error) {
	var dbCompetitions []models.CompetitionModel

	if err := repo.db.Preload("Category").Preload("Weapon").Find(&dbCompetitions).Error; err != nil {
		return nil, err
	}

	competitions := make([]*entities.Competition, len(dbCompetitions))
	for i, dbCompetition := range dbCompetitions {
		competitions[i] = repo.unmarshalCompetition(&dbCompetition)
	}

	return competitions, nil
}

func (repo InMemoryCompetitionRepository) unmarshalCompetition(m *models.CompetitionModel) *entities.Competition {
	competitionCategory := entities.UnmarshalCompetitionCategory(m.Category.ID, m.Category.Name, m.Category.CreatedAt, m.Category.UpdatedAt)

	weapon := entities.UnmarshalWeapon(m.Weapon.ID, m.Weapon.Name, m.Weapon.CreatedAt, m.Weapon.UpdatedAt)

	return entities.UnmarshalCompetition(m.ID, m.CreatedAt, m.UpdatedAt, m.Name, m.OrganizerName, m.FederationName, m.CompetitionType, *competitionCategory, m.Gender, *weapon, m.Date)
}

func (repo InMemoryCompetitionRepository) marshalCompetition(c *entities.Competition) models.CompetitionModel {
	competitionModel := models.CompetitionModel{
		Model:           common.Model{ID: c.ID, CreatedAt: c.CreatedAt, UpdatedAt: c.UpdatedAt},
		Name:            c.Name(),
		OrganizerName:   c.OrganizerName(),
		FederationName:  c.FederationName(),
		CompetitionType: c.CompetitionType(),
		Gender:          c.Gender(),
		Date:            c.Date(),
	}

	return competitionModel
}
