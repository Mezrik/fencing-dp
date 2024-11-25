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
	CreateCompetitor  command.CreateCompetitorHandler
	UpdateCompetitor  command.UpdateCompetitorHandler
	AssignParticipant command.AssignParticipantHandler
	ImportCompetitor  command.ImportCompetitorHandler
}

type Queries struct {
	AllCompetitors  query.AllCompetitorsHandler
	AllParticipants query.AllParticipantsHandler
}

func NewCompetitorService(competitorRepo repositories.CompetitorRepo, clubRepo repositories.ClubRepo, logger *logrus.Entry) Service {

	return Service{
		Commands: Commands{
			CreateCompetitor:  command.NewCreateCompetitorHandler(competitorRepo, clubRepo, logger),
			UpdateCompetitor:  command.NewUpdateCompetitorHandler(competitorRepo, clubRepo, logger),
			AssignParticipant: command.NewAssignParticipantHandler(competitorRepo, logger),
			ImportCompetitor:  command.NewImportCompetitorHandler(competitorRepo, clubRepo, logger),
		},
		Queries: Queries{
			AllCompetitors:  query.NewAllCompetitionsHandler(competitorRepo, clubRepo, logger),
			AllParticipants: query.NewAllParticipantsHandler(competitorRepo, logger),
		},
	}
}
