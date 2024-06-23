package app

import (
	"github.com/Mezrik/fencing-dp/internal/competition/app/command"
	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCompetition command.CreateCompetitionHandler
}

type Queries struct {
	AllCompetitions query.AllCompetitionsHandler
}
