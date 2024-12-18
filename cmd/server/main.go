package main

import (
	"context"
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/common/database"
	"github.com/Mezrik/fencing-dp/internal/common/logger"
	competition "github.com/Mezrik/fencing-dp/internal/competition/app"
	competitor "github.com/Mezrik/fencing-dp/internal/competitor/app"

	competitionRepositories "github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory"
	competitorRepositories "github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory"

	"github.com/Mezrik/fencing-dp/internal/interface/server"
	match "github.com/Mezrik/fencing-dp/internal/match/app"
	matchRepositories "github.com/Mezrik/fencing-dp/internal/match/infrastructure/inmemory"
	"github.com/Mezrik/fencing-dp/migrations"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	logger.Init()

	logger := logrus.NewEntry(logrus.StandardLogger())

	logger.Logger.SetFormatter(&logrus.TextFormatter{})

	db, _ := database.NewConnection(logger, ctx, migrations.SQLiteMigrations)

	competitionRepository := competitionRepositories.NewInMemoryCompetitionRepository(ctx, db)

	competitorRepository := competitorRepositories.NewInMemoryCompetitorRepository(ctx, db)
	clubRepository := competitorRepositories.NewInMemoryClubRepo(ctx, db)

	matchRepository := matchRepositories.NewInMemoryMatchRepo(ctx, db)

	matchService := match.NewMatchService(matchRepository, logger)

	groupRepository := competitionRepositories.NewInMemoryGroupsRepository(ctx, db)

	competitionService := competition.NewCompetitionService(competitionRepository, groupRepository, competitorRepository, logger)
	competitorService := competitor.NewCompetitorService(
		competitorRepository,
		clubRepository,
		logger,
	)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return server.HandlerFromMux(server.NewServer(competitionService, competitorService, matchService), router)
	})
}
