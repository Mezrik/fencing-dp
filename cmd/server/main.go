package main

import (
	"context"
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/common/database/inmemory"
	"github.com/Mezrik/fencing-dp/internal/common/logger"
	"github.com/Mezrik/fencing-dp/internal/competition/app"
	repositories "github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory"
	"github.com/Mezrik/fencing-dp/internal/interface/server"
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

	db, _ := inmemory.NewConnection(logger)

	competitionRepository := repositories.NewInMemoryCompetitionRepository(ctx, db)

	competitionService := app.NewCompetitionService(competitionRepository, logger)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return server.HandlerFromMux(server.NewServer(competitionService), router)
	})
}
