package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type CompetitionDao struct {
	DB *sqlx.DB
}

func (dao *CompetitionDao) Create(competition *models.CompetitionModel) error {
	return nil
}

func (dao *CompetitionDao) FindById(id uuid.UUID) (*models.CompetitionModel, error) {
	return nil, nil
}

func (dao *CompetitionDao) FindAll() ([]*models.CompetitionModel, error) {
	return nil, nil
}

func (dao *CompetitionDao) Update(competition *models.CompetitionModel) error {
	return nil
}

func (dao *CompetitionDao) Delete(id uuid.UUID) error {
	return nil
}
