package server

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s Server) GetCompetitions(w http.ResponseWriter, r *http.Request) {
	competitions, _ := s.competitionService.GetAllCompetitions()

	render.Respond(w, r, competitions)
}
