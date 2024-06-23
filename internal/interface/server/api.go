package server

import (
	"github.com/Mezrik/fencing-dp/internal/competition/app"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	competition app.Application
}

func NewServer(competition app.Application) Server {
	return Server{competition}
}
