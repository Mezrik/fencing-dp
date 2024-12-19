package desktop

import (
	"context"
	"os"

	"github.com/Mezrik/fencing-dp/migrations"

	"github.com/Mezrik/fencing-dp/internal/common/database"
	"github.com/Mezrik/fencing-dp/internal/common/logger"
	competition "github.com/Mezrik/fencing-dp/internal/competition/app"
	competitionRepositories "github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory"
	competitor "github.com/Mezrik/fencing-dp/internal/competitor/app"
	competitorRepositories "github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory"

	match "github.com/Mezrik/fencing-dp/internal/match/app"
	matchRepositories "github.com/Mezrik/fencing-dp/internal/match/infrastructure/inmemory"
	"github.com/sirupsen/logrus"
)

type Admin struct {
	ctx          context.Context
	competitions competition.Service
	competitors  competitor.Service
	matches      match.Service
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

	competitionRepository := competitionRepositories.NewInMemoryCompetitionRepository(a.ctx, db)
	competitorRepository := competitorRepositories.NewInMemoryCompetitorRepository(a.ctx, db)
	clubRepository := competitorRepositories.NewInMemoryClubRepo(a.ctx, db)
	matchRepository := matchRepositories.NewInMemoryMatchRepo(a.ctx, db)
	groupRepository := competitionRepositories.NewInMemoryGroupsRepository(a.ctx, db)

	a.competitions = competition.NewCompetitionService(competitionRepository, groupRepository, competitorRepository, matchRepository, logger)

	a.competitors = competitor.NewCompetitorService(
		competitorRepository,
		clubRepository,
		logger,
	)

	a.matches = match.NewMatchService(matchRepository, logger)
}
