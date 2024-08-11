package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompetitionCategoryDao struct {
	DB *sqlx.DB
}

func (dao *CompetitionCategoryDao) Create(competition *models.CompetitionCategoryModel) error {
	return nil
}

func (dao *CompetitionCategoryDao) FindById(id uuid.UUID) (*models.CompetitionCategoryModel, error) {
	row := dao.DB.QueryRow("SELECT * FROM competition_categories WHERE id = ?", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var competitionCategoryModel models.CompetitionCategoryModel

	if err := row.Scan(&competitionCategoryModel.ID, &competitionCategoryModel.CreatedAt, &competitionCategoryModel.UpdatedAt, &competitionCategoryModel.Name); err != nil {
		return nil, err
	}

	return &competitionCategoryModel, nil
}

func (dao *CompetitionCategoryDao) FindAll() ([]*models.CompetitionCategoryModel, error) {
	return nil, nil
}

func (dao *CompetitionCategoryDao) Update(competition *models.CompetitionCategoryModel) error {
	return nil
}

func (dao *CompetitionCategoryDao) Delete(id uuid.UUID) error {
	return nil
}
