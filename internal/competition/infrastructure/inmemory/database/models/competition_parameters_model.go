package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CompetitionParametersModel struct {
	ID                         uuid.UUID    `db:"id"`
	CreatedAt                  time.Time    `db:"created_at"`
	UpdatedAt                  sql.NullTime `db:"updated_at"`
	ExpectedParticipants       int          `db:"expected_participants"`
	DeploymentType             string       `db:"deployment_type"`
	GroupHits                  int          `db:"group_hits"`
	EliminationHits            int          `db:"elimination_hits"`
	QualificationBasedOnRounds int          `db:"qualification_based_on_rounds"`
}
