package entities

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	competitor        *Competitor
	competition_id    uuid.UUID
	deployment_number int
	points            float64
	starting_position int
}

func NewParticipant(competitor *Competitor, deployment_number int, created_at time.Time, points float64, starting_position int) (*Participant, error) {

	return &Participant{
		competitor:        competitor,
		deployment_number: deployment_number,
		points:            points,
		starting_position: starting_position,
	}, nil
}

func (p Participant) Competitor() *Competitor {
	return p.competitor
}

func (p Participant) CompetitionId() uuid.UUID {
	return p.competition_id
}

func (p Participant) DeploymentNumber() int {
	return p.deployment_number
}

func (p Participant) Points() float64 {
	return p.points
}

func (p Participant) StartingPosition() int {
	return p.starting_position
}
