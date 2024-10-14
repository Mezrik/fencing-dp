package server

import (
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/match/app/query"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (s Server) GetMatchesGroupIdAll(w http.ResponseWriter, r *http.Request, groupId uuid.UUID) {
	matches, _ := s.match.Queries.AllMatches.Handle(r.Context(), query.AllMatches{GroupID: groupId})

	render.Respond(w, r, matches)
}

func (s Server) GetMatchesMatchId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	match, _ := s.match.Queries.GetMatch.Handle(r.Context(), query.GetMatch{ID: id})

	render.Respond(w, r, match)
}
