package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClubDao struct {
	DB *sqlx.DB
}

func (dao *ClubDao) Create(club models.ClubModel) error {
	return nil
}

func (dao *ClubDao) FindAll() ([]*models.ClubModel, error) {
	var clubModels []*models.ClubModel

	err := dao.DB.Select(&clubModels, "SELECT * FROM clubs")

	if err != nil {
		return nil, err
	}

	return clubModels, nil
}

func (dao *ClubDao) FindById(id uuid.UUID) (*models.ClubModel, error) {
	var clubModel models.ClubModel

	err := dao.DB.Get(&clubModel, "SELECT * FROM clubs WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return &clubModel, nil
}
