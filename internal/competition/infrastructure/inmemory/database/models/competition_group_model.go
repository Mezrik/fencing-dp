package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CompetitionGroupModel struct {
	ID            uuid.UUID     `db:"id"`
	Name          string        `db:"name"`
	PisteNumber   sql.NullInt64 `db:"piste_number"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     sql.NullTime  `db:"updated_at"`
	CompetitionId uuid.UUID     `db:"competition_id"`
}
