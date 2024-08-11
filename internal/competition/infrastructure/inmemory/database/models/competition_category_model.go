package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CompetitionCategoryModel struct {
	ID        uuid.UUID    `db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
