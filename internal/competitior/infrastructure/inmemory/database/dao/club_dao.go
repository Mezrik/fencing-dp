package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competitior/infrastructure/inmemory/database/models"
	"github.com/jmoiron/sqlx"
)

type ClubDao struct {
	DB *sqlx.DB
}

func (dao *ClubDao) Create() error {
	return nil
}

func (dao *ClubDao) FindAll() ([]*models.ClubModel, error) {
	return nil, nil
}
