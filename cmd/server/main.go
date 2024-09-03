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

	db, _ := database.NewConnection(logger, ctx, migrations.SQLiteMigrations)

	competitionRepository := competitionRepositories.NewInMemoryCompetitionRepository(ctx, db)

	competitorRepository := competitorRepositories.NewInMemoryCompetitorRepository(ctx, db)
	clubRepository := competitorRepositories.NewInMemoryClubRepo(ctx, db)

	competitorService := competitor.NewCompetitorService(competitorRepository, clubRepository, logger)
	competitionService := competition.NewCompetitionService(competitionRepository, logger)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return server.HandlerFromMux(server.NewServer(competitionService, competitorService), router)
	})
}
