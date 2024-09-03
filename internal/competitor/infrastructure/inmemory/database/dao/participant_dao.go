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
	return nil
}

func (dao *ParticipantDao) FindById(id uuid.UUID) (*models.ParticipantModel, error) {
	return nil, nil
}

func (dao *ParticipantDao) FindAll(competitionId uuid.UUID) ([]*models.ParticipantModel, error) {
	return nil, nil
}
