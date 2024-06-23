package main

import (
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/common/database/inmemory"
	"github.com/Mezrik/fencing-dp/internal/competition"
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

	db, _ := inmemory.NewConnection()

	logger := logrus.NewEntry(logrus.StandardLogger())

	competitionRepository := repositories.NewInMemoryCompetitionRepository(db)

	competitionService := competition.NewCompetitionApp(competitionRepository, logger)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return server.HandlerFromMux(server.NewServer(competitionService), router)
	})
}
