package query

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/google/uuid"
)

type Participant struct {
	Competitor       Competitor    `json:"competitor"`
	CompetitionID    uuid.UUID     `json:"competitionId" ts_type:"UUID"`
	DeploymentNumber int           `json:"deploymentNumber"`
	Points           float64       `json:"points"`
	StartingPosition int           `json:"startingPosition"`
	GroupID          uuid.NullUUID `json:"groupId" ts_type:"UUID"`
}

func ToParticipantQueryFromEntity(c *entities.Participant) *Participant {
	return &Participant{
		Competitor:       *ToCompetitorQueryFromEntity(c.Competitor()),
		CompetitionID:    c.CompetitionId(),
		DeploymentNumber: c.DeploymentNumber(),
		Points:           c.Points(),
		StartingPosition: c.StartingPosition(),
		GroupID:          c.GroupId(),
	}
}

func ToParticipantQueryListFromEntities(c []*entities.Participant) []*Participant {
	participants := make([]*Participant, 0, len(c))

	for _, v := range c {
		participants = append(participants, ToParticipantQueryFromEntity(v))
	}
	return participants
}
