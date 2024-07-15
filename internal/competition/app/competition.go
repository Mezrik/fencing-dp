package app

import (
	"github.com/Mezrik/fencing-dp/internal/competition/app/command"
	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCompetition command.CreateCompetitionHandler
}

type Queries struct {
	AllCompetitions query.AllCompetitionsHandler
	AllCategories   query.AllCategoriesHandler
}

func NewCompetitionService(competitionRepo repositories.CompetitionRepository, logger *logrus.Entry) Service {

	return Service{
		Commands: Commands{
			CreateCompetition: command.NewCreateCompetitionHandler(competitionRepo, logger),
		},
		Queries: Queries{
			AllCompetitions: query.NewAllCompetitionsHandler(competitionRepo, logger),
			AllCategories:   query.NewAllCategoriesHandler(competitionRepo, logger),
		},
	}
}
