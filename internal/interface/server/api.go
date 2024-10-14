package server

import (
	competition "github.com/Mezrik/fencing-dp/internal/competition/app"
	competitor "github.com/Mezrik/fencing-dp/internal/competitor/app"
	match "github.com/Mezrik/fencing-dp/internal/match/app"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	competition competition.Service
	competitor  competitor.Service
	match       match.Service
}

func NewServer(competition competition.Service, competitor competitor.Service, match match.Service) Server {
	return Server{competition, competitor, match}
}
