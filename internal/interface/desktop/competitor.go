package desktop

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/app/command"
	"github.com/Mezrik/fencing-dp/internal/competitor/app/query"
	"github.com/google/uuid"
)

func (a *Admin) GetCompetitors() []*query.Competitor {
	competitors, _ := a.competitors.Queries.AllCompetitors.Handle(a.ctx, query.AllCompetitors{})

	return competitors
}

func (a *Admin) CreateCompetitor(competitor command.CreateCompetitor) error {
	err := a.competitors.Commands.CreateCompetitor.Handle(a.ctx, competitor)

	return err
}

func (a *Admin) GetParticipants(competitionId uuid.UUID) []*query.Participant {
	participants, _ := a.competitors.Queries.AllParticipants.Handle(a.ctx, query.AllParticipants{CompetitionId: competitionId})

	return participants
}

func (a *Admin) AssignCompetitor(competitorId uuid.UUID, competitionId uuid.UUID) error {
	err := a.competitors.Commands.AssignParticipant.Handle(a.ctx, command.AssignParticipant{CompetitionId: competitionId, ParticipantId: competitorId})

	return err
}
