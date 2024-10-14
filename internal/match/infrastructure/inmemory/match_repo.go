package inmemory

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/util"
	"github.com/Mezrik/fencing-dp/internal/match/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/match/infrastructure/inmemory/database/dao"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InMemoryMatchRepo struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewInMemoryMatchRepo(ctx context.Context, db *sqlx.DB) *InMemoryMatchRepo {
	return &InMemoryMatchRepo{db: db}
}

func (repo InMemoryMatchRepo) Create(match *entities.Match) error {
	return nil
}

func (repo InMemoryMatchRepo) FindAll(groupID uuid.UUID) ([]*entities.Match, error) {

	dao := &dao.MatchDao{DB: repo.db}

	matchModels, err := dao.FindAll(groupID)

	if err != nil {
		return nil, err
	}

	matchList := make([]*entities.Match, 0, len(matchModels))

	for _, matchModel := range matchModels {

		matchList = append(matchList, entities.UnmarshalMatch(
			matchModel.ID,
			matchModel.GroupID,
			matchModel.ParticipantOneID,
			matchModel.ParticipantTwoID,
			nil,
			matchModel.CreatedAt,
			util.GetTimePtr(matchModel.UpdatedAt),
		))
	}

	return matchList, nil
}

func (repo InMemoryMatchRepo) FindById(id uuid.UUID) (*entities.Match, error) {
	dao := &dao.MatchDao{DB: repo.db}

	matchModel, err := dao.FindById(id)

	if err != nil {
		return nil, err
	}

	matchList := make([]*entities.MatchState, 0, len(matchModel.State))

	for _, matchStateModel := range matchModel.State {

		match := entities.UnmarshalMatchState(
			matchStateModel.ID,
			matchStateModel.MatchID,
			entities.MatchChangeEnum(matchStateModel.Change),
			matchStateModel.PointTo,
			matchModel.CreatedAt,
			util.GetTimePtr(matchModel.UpdatedAt),
		)

		matchList = append(matchList, match)
	}

	match := entities.UnmarshalMatch(
		matchModel.ID,
		matchModel.GroupID,
		matchModel.ParticipantOneID,
		matchModel.ParticipantTwoID,
		matchList,
		matchModel.CreatedAt,
		util.GetTimePtr(matchModel.UpdatedAt),
	)

	return match, nil
}
