package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompetitionParametersDao struct {
	DB *sqlx.DB
}

func (dao *CompetitionParametersDao) Create(competition *models.CompetitionParametersModel) error {
	query := `
    INSERT INTO competition_parameters (
      id,
      created_at,
      updated_at,
      expected_participants,
      deployment_type,
      group_hits,
      elimination_hits,
      qualification_based_on_rounds
    ) VALUES (
      ?, ?, ?, ?, ?, ?, ?, ?
    )
  `

	_, err := dao.DB.Exec(
		query,
		competition.ID,
		competition.CreatedAt,
		competition.UpdatedAt,
		competition.ExpectedParticipants,
		competition.DeploymentType,
		competition.GroupHits,
		competition.EliminationHits,
		competition.QualificationBasedOnRounds,
	)

	return err
}

func (dao *CompetitionParametersDao) FindById(id uuid.UUID) (*models.CompetitionParametersModel, error) {
	var competitionParametersModel models.CompetitionParametersModel

	err := dao.DB.Get(&competitionParametersModel, "SELECT * FROM competition_parameters WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return &competitionParametersModel, nil
}

func (dao *CompetitionParametersDao) Update(competition *models.CompetitionParametersModel) error {
	query := `
		UPDATE competition
		SET
			expected_participants = :expected_participants,
      deployment_type = :deployment_type,
      group_hits = :group_hits,
      elimination_hits = :elimination_hits,
      qualification_based_on_rounds = :qualification_based_on_rounds
		WHERE id = :id
	`

	_, err := dao.DB.NamedExec(query, competition)

	return err
}
