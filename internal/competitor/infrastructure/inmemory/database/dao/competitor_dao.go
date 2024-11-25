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
	query := `
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
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := c.DB.Exec(
		query,
		competitor.ID,
		competitor.CreatedAt,
		competitor.UpdatedAt,
		competitor.Surname,
		competitor.Firstname,
		competitor.Gender,
		competitor.ClubID,
		competitor.License,
		competitor.LicenseFie,
		competitor.Birthdate,
	)

	return err
}

func (c *CompetitorDao) BatchCreate(competitors []models.CompetitorModel) error {
	query := `
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
		) VALUES (:id, :created_at, :updated_at, :surname, :firstname, :gender, :club_id, :license, :license_fie, :birthdate)
		`

	_, err := c.DB.NamedExec(query, competitors)

	return err
}

func (c *CompetitorDao) FindAll() ([]*models.CompetitorModel, error) {
	query := `
    SELECT
      COALESCE(clb.name, '') "club.name",
      COALESCE(clb.id, '') "club.id",
      c.*
    FROM competitors c
    LEFT JOIN clubs clb ON c.club_id = clb.id
  `

	var competitorModels []*models.CompetitorModel

	err := c.DB.Select(&competitorModels, query)

	if err != nil {
		return nil, err
	}

	for i := range competitorModels {
		if competitorModels[i].ClubID == nil {
			competitorModels[i].Club = nil
		}
	}

	return competitorModels, nil
}

func (c *CompetitorDao) FindById(id uuid.UUID) (*models.CompetitorModel, error) {
	query := `
    SELECT
      clb.name as "club.name",
      clb.id as "club.id",
      c.*
    FROM competitors c
    JOIN clubs clb ON c.club_id = clb.id
    `

	var competitorModel models.CompetitorModel

	err := c.DB.Get(&competitorModel, query, id)

	if err != nil {
		return nil, err
	}

	return &competitorModel, nil
}
