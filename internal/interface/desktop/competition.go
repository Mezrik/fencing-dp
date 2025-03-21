package desktop

import (
	"github.com/Mezrik/fencing-dp/internal/competition/app/command"
	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
	"github.com/google/uuid"
)

func (a *Admin) GetCompetitions() []*query.Competition {
	competitions, _ := a.competitions.Queries.AllCompetitions.Handle(a.ctx, query.AllCompetitions{})

	return competitions
}

func (a *Admin) GetCompetition(id uuid.UUID) *query.Competition {
	competition, _ := a.competitions.Queries.GetCompetition.Handle(a.ctx, query.GetCompetition{ID: id})

	return competition
}

func (a *Admin) CreateCompetition(competition command.CreateCompetition) error {
	err := a.competitions.Commands.CreateCompetition.Handle(a.ctx, competition)

	return err
}

func (a *Admin) GetCompetitionsCategories() []*query.CompetitionCategory {
	categories, _ := a.competitions.Queries.AllCategories.Handle(a.ctx, query.AllCategories{})

	return categories
}

func (a *Admin) GetCompetitionsWeapons() []*query.Weapon {
	weapons, _ := a.competitions.Queries.AllWeapons.Handle(a.ctx, query.AllWeapons{})

	return weapons
}

func (a *Admin) GetCompetitionsGroups(competitionId uuid.UUID) []*query.Group {
	groups, _ := a.competitions.Queries.AllGroups.Handle(a.ctx, query.AllGroups{CompetitionID: competitionId})

	return groups
}

func (a *Admin) GetGroup(id uuid.UUID) *query.Group {
	group, _ := a.competitions.Queries.GetGroup.Handle(a.ctx, query.GetGroup{ID: id})

	return group
}

func (a *Admin) UpdateCompetitionParameters(updateParams command.UpdateCompetitionParameters) error {
	err := a.competitions.Commands.UpdateCompetitionParameters.Handle(a.ctx, updateParams)

	return err
}

func (a *Admin) UpdateCompetition(competition command.UpdateCompetitionCommand) error {
	err := a.competitions.Commands.UpdateCompetition.Handle(a.ctx, competition)

	return err
}

func (a *Admin) InitializeGroups(competitionId uuid.UUID) error {
	err := a.competitions.Commands.InitializeGroups.Handle(a.ctx, command.InitializeGroups{CompetitionID: competitionId})

	return err
}
