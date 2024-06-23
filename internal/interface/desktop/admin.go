package desktop

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/database/inmemory"
	"github.com/Mezrik/fencing-dp/internal/competition"
	"github.com/Mezrik/fencing-dp/internal/competition/app"
	repositories "github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory"
	"github.com/sirupsen/logrus"
)

type Admin struct {
	ctx          context.Context
	competitions app.Application
}

func New() *Admin {
	return &Admin{}
}

func (a *Admin) Startup(ctx context.Context) {
	a.ctx = ctx

	db, _ := inmemory.NewConnection()

	logger := logrus.NewEntry(logrus.StandardLogger())

	competitionRepository := repositories.NewInMemoryCompetitionRepository(db)

	a.competitions = competition.NewCompetitionApp(competitionRepository, logger)
}
