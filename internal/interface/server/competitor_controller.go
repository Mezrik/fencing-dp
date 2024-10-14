package server

import (
	"encoding/json"
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/competitor/app/command"
	"github.com/Mezrik/fencing-dp/internal/competitor/app/query"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (s Server) GetCompetitors(w http.ResponseWriter, r *http.Request) {
	competitors, _ := s.competitor.Queries.AllCompetitors.Handle(r.Context(), query.AllCompetitors{})

	render.Respond(w, r, competitors)
}

func (s Server) PostCompetitors(w http.ResponseWriter, r *http.Request) {
	var competitor command.CreateCompetitor

	if err := json.NewDecoder(r.Body).Decode(&competitor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.competitor.Commands.CreateCompetitor.Handle(r.Context(), competitor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, http.StatusOK)
}

func (s Server) GetCompetitorsAllCompetitionId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	participants, _ := s.competitor.Queries.AllParticipants.Handle(r.Context(), query.AllParticipants{CompetitionId: id})

	render.Respond(w, r, participants)
}
