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

func (s Server) GetCompetitorsCompetitorId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	competitor, _ := s.competitor.Queries.GetCompetitor.Handle(r.Context(), query.GetCompetitor{ID: id})

	render.Respond(w, r, competitor)
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

func (s Server) PutCompetitorsCompetitorId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var competitor command.UpdateCompetitor

	if err := json.NewDecoder(r.Body).Decode(&competitor); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.competitor.Commands.UpdateCompetitor.Handle(r.Context(),
		command.UpdateCompetitor{
			ID:         id,
			Firstname:  competitor.Firstname,
			Surname:    competitor.Surname,
			License:    competitor.License,
			LicenseFie: competitor.LicenseFie,
			Birthdate:  competitor.Birthdate,
			Gender:     string(competitor.Gender),
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.Respond(w, r, http.StatusOK)
}

func (s Server) GetCompetitorsAllCompetitionId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	participants, _ := s.competitor.Queries.AllParticipants.Handle(r.Context(), query.AllParticipants{CompetitionId: id})

	render.Respond(w, r, participants)
}

func (s Server) PostCompetitorsAssignParticipants(w http.ResponseWriter, r *http.Request) {
	var payload command.AssignParticipants

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.competitor.Commands.AssignParticipants.Handle(r.Context(), payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, http.StatusOK)
}

func (s Server) PostCompetitorsImport(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = s.competitor.Commands.ImportCompetitor.Handle(r.Context(), command.ImportCompetitor{File: file})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
