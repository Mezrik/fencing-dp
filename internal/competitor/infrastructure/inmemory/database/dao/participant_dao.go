package dao

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ParticipantDao struct {
	DB *sqlx.DB
}

func (dao *ParticipantDao) Create(participant models.ParticipantModel) error {
	_, err := dao.DB.Exec(` INSERT INTO participating_competitors (
		competitor_id,
		competition_id,
		deployment_number,
		points,
		starting_position
	) VALUES (
		?, ?, ?, ?, ?
	)
	`, participant.CompetitorID, participant.CompetitionID, participant.DeploymentNumber, participant.Points, participant.StartingPosition)

	return err
}

func (c *ParticipantDao) BatchCreate(participants []models.ParticipantModel) error {
	query := `
		INSERT INTO participating_competitors (
			competitor_id,
			competition_id,
			deployment_number,
			points,
			starting_position
		) VALUES (:competitor_id, :competition_id, :deployment_number, :points, :starting_position)
		`

	_, err := c.DB.NamedExec(query, participants)

	return err
}

func (dao *ParticipantDao) FindById(competitionId uuid.UUID, competitorId uuid.UUID) (*models.ParticipantModel, error) {
	var participantModel models.ParticipantModel

	err := dao.DB.Get(&participantModel, `
		SELECT 
			pc.*,
			c.id as "competitor.id",
			c.firstname as "competitor.firstname",
			c.surname as "competitor.surname",
			c.gender as "competitor.gender",
			c.club_id as "competitor.club_id",
			c.license as "competitor.license",
			c.license_fie as "competitor.license_fie",
			c.birthdate as "competitor.birthdate"
		FROM participating_competitors pc
		JOIN competitors c ON pc.competitor_id = c.competitor_id
		WHERE pc.competition_id = ? AND pc.competitor_id = ?
		`, competitionId, competitorId)

	if err != nil {
		return nil, err
	}

	return &participantModel, nil
}

func (dao *ParticipantDao) FindAll(competitionId uuid.UUID) ([]*models.ParticipantModel, error) {
	var participantModels []*models.ParticipantModel

	err := dao.DB.Select(&participantModels, `
		SELECT 
			pc.*,
			c.id as "competitor.id",
			c.firstname as "competitor.firstname",
			c.surname as "competitor.surname",
			c.gender as "competitor.gender",
			c.club_id as "competitor.club_id",
			c.license as "competitor.license",
			c.license_fie as "competitor.license_fie",
			c.birthdate as "competitor.birthdate",
			cl.name as "competitor.club.name",
			cl.created_at as "competitor.club.created_at",
			cl.updated_at as "competitor.club.updated_at",
			cl.id as "competitor.club.id"
		FROM participating_competitors pc
		JOIN competitors c ON pc.competitor_id = c.id
		JOIN clubs cl ON c.club_id = cl.id
		WHERE pc.competition_id = ?
		`, competitionId)

	if err != nil {
		return nil, err
	}

	return participantModels, nil
}
