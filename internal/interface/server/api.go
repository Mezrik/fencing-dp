package server

import (
	competition "github.com/Mezrik/fencing-dp/internal/competition/app"
	competitor "github.com/Mezrik/fencing-dp/internal/competitor/app"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	competition competition.Service
	competitor  competitor.Service
}

func NewServer(competition competition.Service, competitor competitor.Service) Server {
	return Server{competition, competitor}
}
