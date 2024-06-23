package desktop

import (
	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
)

func (a *Admin) GetCompetitions() []*query.Competition {

	competitions, _ := a.competitions.Queries.AllCompetitions.Handle(a.ctx, query.AllCompetitions{})

	return competitions
}
