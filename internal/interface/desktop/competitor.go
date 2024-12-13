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

func (a *Admin) GetCompetitor(id uuid.UUID) *query.Competitor {
	competitor, _ := a.competitors.Queries.GetCompetitor.Handle(a.ctx, query.GetCompetitor{ID: id})

	return competitor
}

func (a *Admin) CreateCompetitor(competitor command.CreateCompetitor) error {
	err := a.competitors.Commands.CreateCompetitor.Handle(a.ctx, competitor)

	return err
}

func (a *Admin) UpdateCompetitor(competitor command.UpdateCompetitor) error {
	err := a.competitors.Commands.UpdateCompetitor.Handle(a.ctx, competitor)

	return err
}

func (a *Admin) GetParticipants(competitionId uuid.UUID) []*query.Participant {
	participants, _ := a.competitors.Queries.AllParticipants.Handle(a.ctx, query.AllParticipants{CompetitionId: competitionId})

	return participants
}

func (a *Admin) AssignCompetitors(competitorIds []uuid.UUID, competitionId uuid.UUID) error {
	err := a.competitors.Commands.AssignParticipants.Handle(a.ctx, command.AssignParticipants{CompetitionId: competitionId, ParticipantIds: competitorIds})

	return err
}
