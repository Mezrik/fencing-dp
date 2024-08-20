package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type WeaponDao struct {
	DB *sqlx.DB
}

func (dao *WeaponDao) Create(competition *models.WeaponModel) error {
	return nil
}

func (dao *WeaponDao) FindById(id uuid.UUID) (*models.WeaponModel, error) {
	row := dao.DB.QueryRow("SELECT * FROM weapons WHERE id = ?", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var weaponModel models.WeaponModel

	if err := row.Scan(&weaponModel.ID, &weaponModel.CreatedAt, &weaponModel.UpdatedAt, &weaponModel.Name); err != nil {
		return nil, err
	}

	return &weaponModel, nil
}

func (dao *WeaponDao) FindAll() ([]*models.WeaponModel, error) {
	var weaponModels []*models.WeaponModel

	err := dao.DB.Select(&weaponModels, "SELECT * FROM weapons")

	if err != nil {
		return nil, err
	}

	return weaponModels, nil
}

func (dao *WeaponDao) Update(competition *models.WeaponModel) error {
	return nil
}

func (dao *WeaponDao) Delete(id uuid.UUID) error {
	return nil
}
