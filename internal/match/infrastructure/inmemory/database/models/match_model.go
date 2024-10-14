package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MatchModel struct {
	ID               uuid.UUID    `db:"id"`
	CreatedAt        time.Time    `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
	GroupID          uuid.UUID    `db:"group_id"`
	ParticipantOneID uuid.UUID    `db:"participant_one_id"`
	ParticipantTwoID uuid.UUID    `db:"participant_two_id"`
}

type MatchModelDetail struct {
	ID               uuid.UUID    `db:"id"`
	CreatedAt        time.Time    `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
	GroupID          uuid.UUID    `db:"group_id"`
	ParticipantOneID uuid.UUID    `db:"participant_one_id"`
	ParticipantTwoID uuid.UUID    `db:"participant_two_id"`
	State            []*MatchStateModel
}
