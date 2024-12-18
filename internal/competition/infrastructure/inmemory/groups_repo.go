package repositories

import (
	"context"
	"database/sql"

	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/dao"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/jmoiron/sqlx"
)

type InMemoryGroupsRepository struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewInMemoryGroupsRepository(ctx context.Context, db *sqlx.DB) *InMemoryGroupsRepository {
	return &InMemoryGroupsRepository{db: db}
}

func (m *InMemoryGroupsRepository) Create(group *entities.Group) error {
	dao := &dao.CompetitionGroupDao{DB: m.db}

	competitionGroupModel, err := m.marshalCompetitionGroup(group)

	if err != nil {
		return err
	}

	return dao.Create(competitionGroupModel)
}

func (m *InMemoryGroupsRepository) marshalCompetitionGroup(group *entities.Group) (*models.CompetitionGroupModel, error) {
	pisteNumber := sql.NullInt64{Valid: true, Int64: 0}

	if group.PisteNumber() != nil {
		pisteNumber = sql.NullInt64{Valid: true, Int64: *group.PisteNumber()}
	}

	var updatedAt sql.NullTime

	if group.UpdatedAt != nil {
		updatedAt = sql.NullTime{Valid: true, Time: *group.UpdatedAt}
	}

	return &models.CompetitionGroupModel{
		ID:            group.ID,
		Name:          group.Name(),
		PisteNumber:   pisteNumber,
		CreatedAt:     group.CreatedAt,
		UpdatedAt:     updatedAt,
		CompetitionId: group.CompetitionID(),
	}, nil
}
