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
	return nil, nil
}

func (dao *ClubDao) FindById(id uuid.UUID) (*models.ClubModel, error) {
	var clubModel models.ClubModel

	err := dao.DB.Get(&clubModel, "SELECT * FROM clubs WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return &clubModel, nil
}
