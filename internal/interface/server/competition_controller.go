package server

import (
	"encoding/json"
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/competition/app/command"
	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
	"github.com/go-chi/render"
)

func (s Server) GetCompetitions(w http.ResponseWriter, r *http.Request) {
	competitions, _ := s.competition.Queries.AllCompetitions.Handle(r.Context(), query.AllCompetitions{})

	render.Respond(w, r, competitions)
}

func (s Server) PostCompetitions(w http.ResponseWriter, r *http.Request) {
	var competition command.CreateCompetition

	if err := json.NewDecoder(r.Body).Decode(&competition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.competition.Commands.CreateCompetition.Handle(r.Context(), competition)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, http.StatusOK)
}

func (s Server) GetCompetitionsCategories(w http.ResponseWriter, r *http.Request) {
	categories, _ := s.competition.Queries.AllCategories.Handle(r.Context(), query.AllCategories{})

	render.Respond(w, r, categories)
}
