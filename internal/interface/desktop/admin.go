package desktop

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/database/inmemory"
	"github.com/Mezrik/fencing-dp/internal/common/logger"
	competition "github.com/Mezrik/fencing-dp/internal/competition/app"
	repositories "github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory"
	"github.com/sirupsen/logrus"
)

type Admin struct {
	ctx          context.Context
	competitions competition.Service
}

func New() *Admin {
	return &Admin{}
}

func (a *Admin) Startup(ctx context.Context) {
	a.ctx = ctx

	logger.Init()

	logger := logrus.NewEntry(logrus.StandardLogger())

	db, _ := inmemory.NewConnection(logger)

	competitionRepository := repositories.NewInMemoryCompetitionRepository(a.ctx, db)

	a.competitions = competition.NewCompetitionService(competitionRepository, logger)
}
