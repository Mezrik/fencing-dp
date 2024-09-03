package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CompetitorModel struct {
	ID         uuid.UUID    `db:"id"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
	Surname    string       `db:"surname"`
	Firstname  string       `db:"firstname"`
	ClubID     uuid.UUID    `db:"club_id"`
	Club       ClubModel    `db:"club"`
	Gender     string       `db:"gender"`
	License    string       `db:"license"`
	LicenseFie *string      `db:"license_fie"`
	Birthdate  time.Time    `db:"birthdate"`
}
