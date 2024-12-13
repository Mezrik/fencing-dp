package app

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/app/command"
	"github.com/Mezrik/fencing-dp/internal/competitor/app/query"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCompetitor   command.CreateCompetitorHandler
	UpdateCompetitor   command.UpdateCompetitorHandler
	AssignParticipants command.AssignParticipantsHandler
	ImportCompetitor   command.ImportCompetitorHandler
}

type Queries struct {
	AllCompetitors  query.AllCompetitorsHandler
	AllParticipants query.AllParticipantsHandler
	GetCompetitor   query.GetCompetitorHandler
}

func NewCompetitorService(competitorRepo repositories.CompetitorRepo, clubRepo repositories.ClubRepo, logger *logrus.Entry) Service {

	return Service{
		Commands: Commands{
			CreateCompetitor:   command.NewCreateCompetitorHandler(competitorRepo, clubRepo, logger),
			UpdateCompetitor:   command.NewUpdateCompetitorHandler(competitorRepo, clubRepo, logger),
			AssignParticipants: command.NewAssignParticipantsHandler(competitorRepo, logger),
			ImportCompetitor:   command.NewImportCompetitorHandler(competitorRepo, clubRepo, logger),
		},
		Queries: Queries{
			AllCompetitors:  query.NewAllCompetitionsHandler(competitorRepo, clubRepo, logger),
			AllParticipants: query.NewAllParticipantsHandler(competitorRepo, logger),
			GetCompetitor:   query.NewGetCompetitorHandler(competitorRepo, clubRepo, logger),
		},
	}
}
