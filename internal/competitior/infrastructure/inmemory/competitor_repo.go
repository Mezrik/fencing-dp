package inmemory

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/util"
	"github.com/Mezrik/fencing-dp/internal/competitior/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competitior/infrastructure/inmemory/database/dao"
	"github.com/Mezrik/fencing-dp/internal/competitior/infrastructure/inmemory/database/models"
	"github.com/jmoiron/sqlx"
)

type InMemoryCompetitorRepository struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewInMemoryCompetitorRepository(ctx context.Context, db *sqlx.DB) *InMemoryCompetitorRepository {

	return &InMemoryCompetitorRepository{db: db}
}

func (repo InMemoryCompetitorRepository) Create(competior *entities.Competitor) error {
	dao := &dao.CompetitorDao{DB: repo.db}

	competitorModel, err := repo.marshalCompetitor(competior)

	if err != nil {
		return err
	}

	return dao.Create(*competitorModel)
}

func (repo InMemoryCompetitorRepository) FindAll() ([]*entities.Competitor, error) {
	dao := &dao.CompetitorDao{DB: repo.db}

	competitorModels, err := dao.FindAll()

	if err != nil {
		return nil, err
	}

	competitors := make([]*entities.Competitor, 0, len(competitorModels))

	for _, c := range competitorModels {
		club := entities.UnmarshalClub(c.Club.ID, c.Club.Name, c.Club.CreatedAt, util.GetTimePtr(c.Club.UpdatedAt))

		competitors = append(
			competitors,
			entities.UnmarshalCompetitor(c.ID, c.Firstname, c.Surname, entities.GenderEnum(c.Gender), *club, c.License, c.LicenseFie, c.Birthdate, c.CreatedAt, util.GetTimePtr(c.UpdatedAt)),
		)

	}
	return competitors, nil
}

func (repo InMemoryCompetitorRepository) marshalCompetitor(c *entities.Competitor) (*models.CompetitorModel, error) {
	competitorModel := &models.CompetitorModel{
		ID:        c.ID,
		Surname:   c.Surname(),
		Firstname: c.FirstName(),
		Gender:    string(c.Gender()),
		ClubID:    c.Club().ID,
		License:   c.License(),
		Birthdate: c.Birthdate(),
	}
	return competitorModel, nil
}
