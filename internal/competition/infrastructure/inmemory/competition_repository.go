package repositories

import (
	"context"
	"database/sql"

	models "github.com/Mezrik/fencing-dp/internal/common/database/generated"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type InMemoryCompetitionRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewInMemoryCompetitionRepository(ctx context.Context, db *sql.DB) repositories.CompetitionRepository {
	return &InMemoryCompetitionRepository{db: db}
}

func (repo InMemoryCompetitionRepository) Create(competition *entities.Competition) error {
	dbCompetition := repo.marshalCompetition(competition)

	return dbCompetition.Insert(repo.ctx, repo.db, boil.Infer())
}

func (repo InMemoryCompetitionRepository) FindById(id uuid.UUID) (*entities.Competition, error) {
	marshalledId, err := id.MarshalBinary()

	if err != nil {
		return nil, err
	}

	dbCompetition, err := models.FindCompetition(repo.ctx, repo.db, marshalledId)

	if err != nil {
		return nil, err
	}

	return repo.unmarshalCompetition(dbCompetition), nil
}

func (repo InMemoryCompetitionRepository) FindAll() ([]*entities.Competition, error) {
	dbCompetitions, err := models.Competitions().All(repo.ctx, repo.db)

	if err != nil {
		return nil, err
	}

	competitions := make([]*entities.Competition, len(dbCompetitions))
	for i, dbCompetition := range dbCompetitions {
		competitions[i] = repo.unmarshalCompetition(dbCompetition)
	}

	return competitions, nil
}

func (repo InMemoryCompetitionRepository) unmarshalCompetition(m *models.Competition) *entities.Competition {
	competitionCategory := entities.UnmarshalCompetitionCategory(uuid.UUID(m.R.Category.ID), m.R.Category.Name, m.R.Category.CreatedAt, m.R.Category.UpdatedAt.Time)

	weapon := entities.UnmarshalWeapon(uuid.UUID(m.R.Weapon.ID), m.R.Weapon.Name, m.R.Weapon.CreatedAt, m.R.Weapon.UpdatedAt.Time)

	return entities.UnmarshalCompetition(uuid.UUID(m.ID), m.CreatedAt, m.UpdatedAt.Time, m.Name, m.OrganizerName, m.FederationName, entities.CompetitionTypeEnum(m.CompetitionType), *competitionCategory, entities.GenderEnum(m.Gender), *weapon, m.Date)
}

func (repo InMemoryCompetitionRepository) marshalCompetition(c *entities.Competition) models.Competition {
	competitionModel := models.Competition{
		ID:              c.ID[:],
		Name:            c.Name(),
		OrganizerName:   c.OrganizerName(),
		FederationName:  c.FederationName(),
		CompetitionType: int64(c.CompetitionType()),
		Gender:          int64(c.Gender()),
		Date:            c.Date(),
	}

	return competitionModel
}
