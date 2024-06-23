package repositories

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/mappers"
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

func (repo *InMemoryCompetitionRepository) Create(competition *entities.Competition) error {
	dbCompetition := mappers.ToCompetitionModel(competition)

	return repo.db.Create(&dbCompetition).Error
}

func (repo *InMemoryCompetitionRepository) FindById(id uuid.UUID) (*entities.Competition, error) {
	var dbCompetition models.CompetitionModel

	if err := repo.db.Preload("Category").Preload("Weapon").First(&dbCompetition, id).Error; err != nil {
		return nil, err
	}

	return mappers.ToCompetitionEntity(&dbCompetition), nil
}

func (repo *InMemoryCompetitionRepository) FindAll() ([]*entities.Competition, error) {
	var dbCompetitions []models.CompetitionModel

	if err := repo.db.Preload("Category").Preload("Weapon").Find(&dbCompetitions).Error; err != nil {
		return nil, err
	}

	competitions := make([]*entities.Competition, len(dbCompetitions))
	for i, dbCompetition := range dbCompetitions {
		competitions[i] = mappers.ToCompetitionEntity(&dbCompetition)
	}

	return competitions, nil
}
