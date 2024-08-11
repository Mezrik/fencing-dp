package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompetitionDao struct {
	DB *sqlx.DB
}

func (dao *CompetitionDao) Create(competition *models.CompetitionModel) error {
	_, err := dao.DB.Exec(`INSERT INTO competitions (id, created_at, updated_at, weapon_id, category_id, name, organizer_name, federation_name, competition_type, gender, date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, competition.ID, competition.CreatedAt, competition.UpdatedAt, competition.WeaponID, competition.CategoryID, competition.Name, competition.OrganizerName, competition.FederationName, competition.CompetitionType, competition.Gender, competition.Date)

	if err != nil {
		return err
	}

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
