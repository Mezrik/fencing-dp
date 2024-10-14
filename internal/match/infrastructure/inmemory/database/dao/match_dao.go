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
	return nil
}

func (dao *MatchDao) FindAll(groupID uuid.UUID) ([]*models.MatchModel, error) {
	var matchModels []*models.MatchModel

	err := dao.DB.Select(&matchModels, "SELECT * FROM matches WHERE group_id = ?", groupID)

	if err != nil {
		return nil, err
	}

	return matchModels, nil
}

func (dao *MatchDao) FindById(id uuid.UUID) (*models.MatchModelDetail, error) {
	var matchModel models.MatchModelDetail

	err := dao.DB.Get(&matchModel, `
    SELECT 
      m.*,
      ms.id as "state.id",
      ms.match_id as "state.match_id",
      ms.change as "state.change",
      ms.point_to as "state.point_to"
    FROM matches m 
    JOIN match_states ms ON m.id = ms.match_id
    WHERE id = ?
  `, id)

	if err != nil {
		return nil, err
	}

	return &matchModel, nil
}
