package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompetitionGroupDao struct {
	DB *sqlx.DB
}

func (dao *CompetitionGroupDao) Create(competitionGroup *models.CompetitionGroupModel) error {
	query := `
    INSERT INTO competition_groups (
      id,
      created_at,
      updated_at,
      name,
      piste_number,
      competition_id
    ) VALUES (?, ?, ?, ?, ?, ?)
  `

	_, err := dao.DB.Exec(
		query,
		competitionGroup.ID,
		competitionGroup.CreatedAt,
		competitionGroup.UpdatedAt,
		competitionGroup.Name,
		competitionGroup.PisteNumber,
		competitionGroup.CompetitionId,
	)

	return err
}

func (dao *CompetitionGroupDao) FindAll(competitionId uuid.UUID) ([]*models.CompetitionGroupModel, error) {
	var competitionGroupModels []*models.CompetitionGroupModel

	err := dao.DB.Select(&competitionGroupModels, "SELECT * FROM competition_groups WHERE competition_id = ?", competitionId)

	if err != nil {
		return nil, err
	}

	return competitionGroupModels, nil
}

func (dao *CompetitionGroupDao) FindById(id uuid.UUID) (*models.CompetitionGroupModel, error) {
	var competitionGroupModel models.CompetitionGroupModel

	err := dao.DB.Get(&competitionGroupModel, "SELECT * FROM competition_groups WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return &competitionGroupModel, nil
}
