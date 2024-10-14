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
	GetCompetition  query.GetCompetitionHandler
	AllCategories   query.AllCategoriesHandler
	AllWeapons      query.AllWeaponsHandler
	AllGroups       query.AllGroupsHandler
	GetGroup        query.GetGroupHandler
}

func NewCompetitionService(competitionRepo repositories.CompetitionRepository, logger *logrus.Entry) Service {

	return Service{
		Commands: Commands{
			CreateCompetition: command.NewCreateCompetitionHandler(competitionRepo, logger),
		},
		Queries: Queries{
			AllCompetitions: query.NewAllCompetitionsHandler(competitionRepo, logger),
			GetCompetition:  query.NewGetCompetitionHandler(competitionRepo, logger),
			AllCategories:   query.NewAllCategoriesHandler(competitionRepo, logger),
			AllWeapons:      query.NewAllWeaponsHandler(competitionRepo, logger),
			AllGroups:       query.NewAllGroupsHandler(competitionRepo, logger),
			GetGroup:        query.NewGetGroupHandler(competitionRepo, logger),
		},
	}
}
