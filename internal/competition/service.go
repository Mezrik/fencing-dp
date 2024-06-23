package competition

import (
	"github.com/Mezrik/fencing-dp/internal/competition/app"
	"github.com/Mezrik/fencing-dp/internal/competition/app/command"
	"github.com/Mezrik/fencing-dp/internal/competition/app/query"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
)

func NewCompetitionApp(competitionRepo repositories.CompetitionRepository) app.Application {

	return app.Application{
		Commands: app.Commands{
			CreateCompetition: command.NewCreateTrainingHandler(competitionRepo),
		},
		Queries: app.Queries{
			AllCompetitions: query.NewAllCompetitionsHandler(competitionRepo),
		},
	}
}
