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
	query := `
		INSERT INTO competitions (
			id,
			created_at,
			updated_at,
			weapon_id,
			category_id,
			name,
			organizer_name,
			federation_name,
			competition_type,
			gender,
			date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := dao.DB.Exec(
		query,
		competition.ID,
		competition.CreatedAt,
		competition.UpdatedAt,
		competition.WeaponID,
		competition.CategoryID,
		competition.Name,
		competition.OrganizerName,
		competition.FederationName,
		competition.CompetitionType,
		competition.Gender,
		competition.Date,
	)

	return err
}

func (dao *CompetitionDao) FindById(id uuid.UUID) (*models.CompetitionModel, error) {
	var competition models.CompetitionModel

	err := dao.DB.Get(&competition, `
		SELECT 
				w.id as "weapon.id", 
				w.name as "weapon.name", 
				w.created_at as "weapon.created_at", 
				w.updated_at as "weapon.updated_at", 
				cc.id as "category.id", 
				cc.name as "category.name", 
				cc.created_at as "category.created_at", 
				cc.updated_at as "category.updated_at", 
				cp.id as "parameters.id", 
				cp.created_at as "parameters.created_at",
				cp.updated_at as "parameters.updated_at",
				cp.expected_participants as "parameters.expected_participants",
				cp.deployment_type as "parameters.deployment_type",
				cp.group_hits as "parameters.group_hits",
				cp.elimination_hits as "parameters.elimination_hits",
				cp.qualification_based_on_rounds as "parameters.qualification_based_on_rounds",
				c.*
			FROM competitions c
			JOIN weapons w ON c.weapon_id = w.id
      JOIN competition_categories cc ON c.category_id = cc.id
			JOIN competition_parameters cp ON c.competition_parameters_id = cp.id
			WHERE c.id = ?
		`, id)

	if err != nil {
		return nil, err
	}

	return &competition, nil
}

func (dao *CompetitionDao) FindAll() ([]*models.CompetitionModel, error) {
	query := `
        SELECT 
					w.id as "weapon.id", 
					w.name as "weapon.name", 
					w.created_at as "weapon.created_at", 
					w.updated_at as "weapon.updated_at", 
					cc.id as "category.id", 
					cc.name as "category.name", 
					cc.created_at as "category.created_at", 
					cc.updated_at as "category.updated_at", 
					c.*
        FROM competitions c
        JOIN weapons w ON c.weapon_id = w.id
        JOIN competition_categories cc ON c.category_id = cc.id
    `

	// Define a slice to hold the results
	var competitions []*models.CompetitionModel

	// Use sqlx to execute the query and map the rows directly to structs
	err := dao.DB.Select(&competitions, query)

	if err != nil {
		return nil, err
	}

	return competitions, nil
}

func (dao *CompetitionDao) Update(competition *models.CompetitionModel) error {
	query := `
		UPDATE competition
		SET
			updated_at = :updated_at,
			weapon_id = :weapon_id,
			category_id = :category_id,
			name = :name,
			organizer_name = :organizer_name,
			federation_name = :federation_name,
			competition_type = :competition_type,
			gender = :gender,
			date = :date
			parameters_id = :parameters_id
		WHERE id = :id
	`

	_, err := dao.DB.NamedExec(query, competition)

	return err
}

func (dao *CompetitionDao) Delete(id uuid.UUID) error {
	return nil
}
