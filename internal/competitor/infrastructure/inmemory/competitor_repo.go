package inmemory

import (
	"context"
	"database/sql"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common/util"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/dao"
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
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

func (repo InMemoryCompetitorRepository) BatchCreate(competitors []*entities.Competitor) error {
	dao := &dao.CompetitorDao{DB: repo.db}

	competitorModels := make([]models.CompetitorModel, 0, len(competitors))

	for _, c := range competitors {
		competitorModel, err := repo.marshalCompetitor(c)

		if err != nil {
			return err
		}

		competitorModels = append(competitorModels, *competitorModel)
	}

	return dao.BatchCreate(competitorModels)
}

func (repo InMemoryCompetitorRepository) Update(competitor *entities.Competitor) error {
	dao := &dao.CompetitorDao{DB: repo.db}

	competitorModel, err := repo.marshalCompetitor(competitor)

	if err != nil {
		return err
	}

	return dao.Update(*competitorModel)
}

func (repo InMemoryCompetitorRepository) FindAll() ([]*entities.Competitor, error) {
	dao := &dao.CompetitorDao{DB: repo.db}

	competitorModels, err := dao.FindAll()

	if err != nil {
		return nil, err
	}

	competitors := make([]*entities.Competitor, 0, len(competitorModels))

	for _, c := range competitorModels {
		var club *entities.Club

		if c.Club != nil {
			club = entities.UnmarshalClub(c.Club.ID, c.Club.Name, c.Club.CreatedAt, util.GetTimePtr(c.Club.UpdatedAt))
		}

		competitors = append(
			competitors,
			entities.UnmarshalCompetitor(c.ID, c.Firstname, c.Surname, entities.GenderEnum(c.Gender), club, c.License, c.LicenseFie, util.GetTimePtr(c.Birthdate), c.CreatedAt, util.GetTimePtr(c.UpdatedAt)),
		)

	}
	return competitors, nil
}

func (repo InMemoryCompetitorRepository) FindById(id uuid.UUID) (*entities.Competitor, error) {
	dao := &dao.CompetitorDao{DB: repo.db}

	competitorModel, err := dao.FindById(id)

	if err != nil {
		return nil, err
	}

	club := entities.UnmarshalClub(*competitorModel.ClubID, competitorModel.Club.Name, competitorModel.Club.CreatedAt, util.GetTimePtr(competitorModel.Club.UpdatedAt))

	competitor := entities.UnmarshalCompetitor(
		competitorModel.ID,
		competitorModel.Firstname,
		competitorModel.Surname,
		entities.GenderEnum(competitorModel.Gender),
		club,
		competitorModel.License,
		competitorModel.LicenseFie,
		util.GetTimePtr(competitorModel.Birthdate),
		competitorModel.CreatedAt,
		util.GetTimePtr(competitorModel.UpdatedAt),
	)

	return competitor, nil
}

func (repo InMemoryCompetitorRepository) FindAllByCompetitionId(competitionId uuid.UUID) ([]*entities.Participant, error) {

	participantDao := &dao.ParticipantDao{DB: repo.db}

	participantModels, err := participantDao.FindAll(competitionId)

	if err != nil {
		return nil, err
	}

	participants := make([]*entities.Participant, 0, len(participantModels))

	for _, p := range participantModels {

		club := entities.UnmarshalClub(*p.Competitor.ClubID, p.Competitor.Club.Name, p.Competitor.Club.CreatedAt, util.GetTimePtr(p.Competitor.Club.UpdatedAt))

		participants = append(
			participants,
			entities.UnmarshalParticipant(
				p.CompetitionID,
				entities.UnmarshalCompetitor(
					p.Competitor.ID,
					p.Competitor.Firstname,
					p.Competitor.Surname,
					entities.GenderEnum(p.Competitor.Gender),
					club,
					p.Competitor.License,
					p.Competitor.LicenseFie,
					util.GetTimePtr(p.Competitor.Birthdate),
					p.Competitor.CreatedAt,
					util.GetTimePtr(p.Competitor.UpdatedAt),
				),
				p.DeploymentNumber,
				p.Points,
				p.StartingPosition,
				p.GroupID,
			),
		)
	}

	return participants, nil
}

func (repo InMemoryCompetitorRepository) AssignCompetitor(competitorId uuid.UUID, competitionId uuid.UUID) error {
	participantDao := &dao.ParticipantDao{DB: repo.db}

	competitorDao := &dao.CompetitorDao{DB: repo.db}

	competitorModel, err := competitorDao.FindById(competitorId)

	if err != nil {
		return err
	}

	club := entities.UnmarshalClub(*competitorModel.ClubID, competitorModel.Club.Name, competitorModel.Club.CreatedAt, util.GetTimePtr(competitorModel.Club.UpdatedAt))

	participant, err := entities.NewParticipant(
		entities.UnmarshalCompetitor(
			competitorModel.ID,
			competitorModel.Firstname,
			competitorModel.Surname,
			entities.GenderEnum(competitorModel.Gender),
			club,
			competitorModel.License,
			competitorModel.LicenseFie,
			util.GetTimePtr(competitorModel.Birthdate),
			competitorModel.CreatedAt,
			&competitorModel.UpdatedAt.Time,
		),
		competitionId,
	)

	if err != nil {
		return err
	}

	participantModel, err := repo.marshalParticipant(participant)

	if err != nil {
		return err
	}

	return participantDao.Create(participantModel)
}

func (repo InMemoryCompetitorRepository) marshalCompetitor(c *entities.Competitor) (*models.CompetitorModel, error) {
	var clubID *uuid.UUID

	if c.Club() != nil {
		clubID = &c.Club().ID
	}

	var updatedAt time.Time

	if c.UpdatedAt != nil {
		updatedAt = *c.UpdatedAt
	}

	competitorModel := &models.CompetitorModel{
		ID:        c.ID,
		Surname:   c.Surname(),
		Firstname: c.FirstName(),
		Gender:    string(c.Gender()),
		ClubID:    clubID,
		License:   c.License(),
		Birthdate: sql.NullTime{
			Valid: !c.Birthdate().IsZero(),
			Time:  *c.Birthdate(),
		},
		CreatedAt: c.CreatedAt,
		UpdatedAt: sql.NullTime{
			Valid: !updatedAt.IsZero(),
			Time:  updatedAt,
		},
	}
	return competitorModel, nil
}

func (repo InMemoryCompetitorRepository) marshalParticipant(c *entities.Participant) (*models.ParticipantModel, error) {

	participantModel := &models.ParticipantModel{
		CompetitionID:    c.CompetitionId(),
		CompetitorID:     c.Competitor().ID,
		DeploymentNumber: c.DeploymentNumber(),
		Points:           c.Points(),
		StartingPosition: c.StartingPosition(),
	}

	return participantModel, nil
}
