package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/jmoiron/sqlx"
)

type CompetitionGroupRoundDao struct {
	DB *sqlx.DB
}

func (dao *CompetitionGroupRoundDao) Create(competitionGroupRound *models.CompetitionGroupRoundModel) error {
	query := `
    INSERT INTO competition_group_rounds (
      id,
      created_at,
      updated_at,
      competition_id,
      number,
      participants_starting_count,
      shift_criteria,
      number_of_advancers,
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := dao.DB.Exec(
		query,
		competitionGroupRound.ID,
		competitionGroupRound.CreatedAt,
		competitionGroupRound.UpdatedAt,
		competitionGroupRound.CompetitionID,
		competitionGroupRound.Number,
		competitionGroupRound.ParticipantStartingCount,
		competitionGroupRound.ShiftCriteria,
		competitionGroupRound.NumberOfAdvancers,
	)

	return err
}
