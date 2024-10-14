package query

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/google/uuid"
)

type Participant struct {
	Competitor       Competitor
	CompetitionID    uuid.UUID
	DeploymentNumber int
	Points           float64
	StartingPosition int
}

func ToParticipantQueryFromEntity(c *entities.Participant) *Participant {
	return &Participant{
		Competitor:       *ToCompetitorQueryFromEntity(c.Competitor()),
		CompetitionID:    c.CompetitionId(),
		DeploymentNumber: c.DeploymentNumber(),
		Points:           c.Points(),
		StartingPosition: c.StartingPosition(),
	}
}

func ToParticipantQueryListFromEntities(c []*entities.Participant) []*Participant {
	participants := make([]*Participant, 0, len(c))

	for _, v := range c {
		participants = append(participants, ToParticipantQueryFromEntity(v))
	}
	return participants
}
