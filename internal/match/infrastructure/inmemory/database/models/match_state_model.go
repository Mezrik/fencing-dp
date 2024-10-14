package models

import (
	"github.com/google/uuid"
)

type MatchStateModel struct {
	ID      uuid.UUID     `db:"id"`
	MatchID uuid.UUID     `db:"match_id"`
	Change  string        `db:"change"`
	PointTo uuid.NullUUID `db:"point_to"`
}
