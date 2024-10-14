package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MatchStateModel struct {
	ID        uuid.UUID     `db:"id"`
	MatchID   uuid.UUID     `db:"match_id"`
	Change    string        `db:"change"`
	PointTo   uuid.NullUUID `db:"point_to"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt sql.NullTime  `db:"updated_at"`
}
