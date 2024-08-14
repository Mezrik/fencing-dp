package desktop

import (
	"context"
	"os"

	"github.com/Mezrik/fencing-dp/migrations"

	"github.com/Mezrik/fencing-dp/internal/common/database"
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

	f, err := os.OpenFile("output.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logger.WithError(err).Error("failed to open log file")
	}
	logrus.SetOutput(f)

	db, _ := database.NewConnection(logger, a.ctx, migrations.SQLiteMigrations)

	competitionRepository := repositories.NewInMemoryCompetitionRepository(a.ctx, db)

	a.competitions = competition.NewCompetitionService(competitionRepository, logger)
}
