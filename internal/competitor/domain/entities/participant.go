package entities

import (
	"github.com/google/uuid"
)

type Participant struct {
	competitor        *Competitor
	competition_id    uuid.UUID
	deployment_number int
	points            float64
	starting_position int
}

func NewParticipant(competitor *Competitor, competition_id uuid.UUID) (*Participant, error) {

	return &Participant{
		competition_id:    competition_id,
		competitor:        competitor,
		deployment_number: 0,
		points:            0,
		starting_position: 0,
	}, nil
}

func UnmarshalParticipant(
	competition_id uuid.UUID,
	competitor *Competitor,
	deployment_number int,
	points float64,
	starting_position int,
) *Participant {

	return &Participant{
		competitor:        competitor,
		competition_id:    competition_id,
		deployment_number: deployment_number,
		points:            points,
		starting_position: starting_position,
	}
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
