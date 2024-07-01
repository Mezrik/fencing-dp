package server

import (
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
	"github.com/go-chi/render"
)

func (s Server) GetCompetitions(w http.ResponseWriter, r *http.Request) {
	competitions, _ := s.competition.Queries.AllCompetitions.Handle(r.Context(), query.AllCompetitions{})

	render.Respond(w, r, competitions)
}

func (s Server) PostCompetitions(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusNotImplemented)
}
