package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompetitorDao struct {
	DB *sqlx.DB
}

func (c *CompetitorDao) Create(competitor models.CompetitorModel) error {
	_, err := c.DB.NamedExec(`
    INSERT INTO competitors (
      id, 
      created_at, 
      updated_at, 
      surname, 
      firstname, 
      gender, 
      club_id, 
      license, 
      license_fie, 
      birthdate
    ) VALUES (
      :id, 
      :created_at, 
      :updated_at, 
      :surname, 
      :firstname, 
      :gender, 
      :club_id, 
      :license, 
      :license_fie, 
      :birthdate
    )`, competitor)

	return err
}

func (c *CompetitorDao) FindAll() ([]*models.CompetitorModel, error) {
	query := `
    SELECT
      club.name as "club.name",
      club.id as "club.id",
      c.*
    FROM competitors c
    JOIN club ON competitors.club_id = club.id
  `

	var competitorModels []*models.CompetitorModel

	err := c.DB.Select(&competitorModels, query)

	if err != nil {
		return nil, err
	}

	return competitorModels, nil
}

func (c *CompetitorDao) FindById(id uuid.UUID) (*models.CompetitorModel, error) {
	return nil, nil
}
