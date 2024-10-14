package desktop

import (
	"github.com/Mezrik/fencing-dp/internal/match/app/query"
	"github.com/google/uuid"
)

func (a *Admin) GetMatches(groupID uuid.UUID) []*query.Match {
	matches, _ := a.matches.Queries.AllMatches.Handle(a.ctx, query.AllMatches{GroupID: groupID})

	return matches
}

func (a *Admin) GetMatch(id uuid.UUID) *query.MatchDetail {
	match, _ := a.matches.Queries.GetMatch.Handle(a.ctx, query.GetMatch{ID: id})

	return match
}
