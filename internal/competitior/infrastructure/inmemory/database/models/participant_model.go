package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ParticipantModel struct {
	ID            uuid.UUID    `db:"id"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
	CompetitorID  uuid.UUID    `db:"competitor_id"`
	CompetitionID uuid.UUID    `db:"competition_id"`

	Competitor       CompetitorModel `db:"competitor"`
	DeploymentNumber int             `db:"deployment_number"`
	Points           float64         `db:"points"`
	StartingPosition int             `db:"starting_position"`
}
