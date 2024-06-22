package server

import "github.com/Mezrik/fencing-dp/internal/app/interfaces"

var _ ServerInterface = (*Server)(nil)

type Server struct {
	competitionService interfaces.CompetitionService
}

func NewServer(competitionService interfaces.CompetitionService) Server {
	return Server{competitionService}
}
