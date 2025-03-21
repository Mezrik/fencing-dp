package dao

import (
	"github.com/Mezrik/fencing-dp/internal/match/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MatchDao struct {
	DB *sqlx.DB
}

func (dao *MatchDao) Create(match models.MatchModel) error {

	_, err := dao.DB.NamedExec(`
		INSERT INTO competition_matches
		(
			id,
			competition_group_id,
			participant_one_id,
			participant_two_id,
			created_at,
			updated_at
		)
		VALUES
		(
			:id,
			:competition_group_id,
			:participant_one_id,
			:participant_two_id,
			:created_at,
			:updated_at
		)
	`, match)

	return err
}

func (dao *MatchDao) FindAll(groupID uuid.UUID) ([]*models.MatchModel, error) {
	var matchModels []*models.MatchModel

	err := dao.DB.Select(&matchModels, "SELECT * FROM competition_matches WHERE competition_group_id = ?", groupID)

	if err != nil {
		return nil, err
	}

	return matchModels, nil
}

func (dao *MatchDao) FindById(id uuid.UUID) (*models.MatchModelDetail, error) {
	var matchState []*models.MatchStateModel
	var matchModel models.MatchModelDetail

	errMatchState := dao.DB.Select(&matchState, `
		SELECT * 
		FROM match_states ms
		WHERE ms.match_id = ?
	`, id)

	if errMatchState != nil {
		return nil, errMatchState
	}

	err := dao.DB.Get(&matchModel, `
    SELECT *
    FROM competition_matches
    WHERE id = ?
  `, id)

	if err != nil {
		return nil, err
	}

	matchModel.State = matchState

	return &matchModel, nil
}
