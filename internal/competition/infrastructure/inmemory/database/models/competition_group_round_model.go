package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CompetitionGroupRoundModel struct {
	ID                       uuid.UUID    `db:"id"`
	CreatedAt                time.Time    `db:"created_at"`
	UpdatedAt                sql.NullTime `db:"updated_at"`
	Number                   int          `db:"number"`
	CompetitionID            uuid.UUID    `db:"competition_id"`
	ParticipantStartingCount int          `db:"participant_starting_count"`
	NumberOfGroups           int          `db:"number_of_groups"`
	ParticipantsInGroups     int          `db:"participants_in_groups"`
	ShiftCriteria            string       `db:"shift_criteria"`
	NumberOfAdvancers        int          `db:"number_of_advancers"`
}
