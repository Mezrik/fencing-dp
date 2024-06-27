package server

import (
	competition "github.com/Mezrik/fencing-dp/internal/competition/app"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	competition competition.Service
}

func NewServer(competition competition.Service) Server {
	return Server{competition}
}
