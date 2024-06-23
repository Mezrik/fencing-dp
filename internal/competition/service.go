package competition

import (
	"github.com/Mezrik/fencing-dp/internal/competition/app"
	"github.com/Mezrik/fencing-dp/internal/competition/app/command"
	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/sirupsen/logrus"
)

func NewCompetitionApp(competitionRepo repositories.CompetitionRepository, logger *logrus.Entry) app.Application {

	return app.Application{
		Commands: app.Commands{
			CreateCompetition: command.NewCreateCompetitionHandler(competitionRepo, logger),
		},
		Queries: app.Queries{
			AllCompetitions: query.NewAllCompetitionsHandler(competitionRepo, logger),
		},
	}
}
