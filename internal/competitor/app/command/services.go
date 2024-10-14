package command

import (
	"context"

	"github.com/google/uuid"
)

type GetCompetition struct {
	ID uuid.UUID
}
type CompetitionService interface {
	GetCompetition(ctx context.Context, query GetCompetition) (interface{}, error)
}
