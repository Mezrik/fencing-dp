package app

import "github.com/Mezrik/fencing-dp/internal/app/command"

type Application struct {
}

type Command struct {
	CreateCompetition command.CreateCompetitionHandler
}
