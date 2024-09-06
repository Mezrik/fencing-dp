package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ParticipantDao struct {
	DB *sqlx.DB
}

func (dao *ParticipantDao) Create(participant *models.ParticipantModel) error {
	_, err := dao.DB.Exec(` INSERT INTO participants (
		id,
		created_at,
		competitor_id,
		competition_id,
		deployment_number,
		points,
		starting_position
	) VALUES (
		:id,
		:created_at,
		:competitor_id,
		:competition_id,
		:deployment_number,
		:points,
		:starting_position
	)
	`, participant)

	return err
}

func (dao *ParticipantDao) FindById(id uuid.UUID) (*models.ParticipantModel, error) {
	var participantModel models.ParticipantModel

	err := dao.DB.Get(&participantModel, "SELECT * FROM participants WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return &participantModel, nil
}

func (dao *ParticipantDao) FindAll(competitionId uuid.UUID) ([]*models.ParticipantModel, error) {
	var participantModels []*models.ParticipantModel

	err := dao.DB.Select(&participantModels, "SELECT * FROM participants WHERE competition_id = ?", competitionId)

	if err != nil {
		return nil, err
	}

	return participantModels, nil
}
