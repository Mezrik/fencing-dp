package entities

import "time"

type Participant struct {
	competitor        *Competitor
	club              *Club
	deployment_number int
	points            float64
	starting_position int
}

func NewParticipant(competitor *Competitor, club *Club, deployment_number int, created_at time.Time, points float64, starting_position int) (*Participant, error) {

	return &Participant{
		competitor:        competitor,
		club:              club,
		deployment_number: deployment_number,
		points:            points,
		starting_position: starting_position,
	}, nil
}

func (p *Participant) Competitor() *Competitor {
	return p.competitor
}

func (p *Participant) Club() *Club {
	return p.club
}

func (p *Participant) DeploymentNumber() int {
	return p.deployment_number
}

func (p *Participant) Points() float64 {
	return p.points
}
