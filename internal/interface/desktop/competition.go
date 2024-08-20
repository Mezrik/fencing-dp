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
