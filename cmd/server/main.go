package main

import (
	"net/http"

	"github.com/Mezrik/fencing-dp/internal/app/services"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/database/inmemory"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/database/inmemory/repositories"
	"github.com/Mezrik/fencing-dp/internal/interface/server"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, _ := inmemory.NewConnection()

	competitionRepository := repositories.NewInMemoryCompetitionRepository(db)

	competitionService := services.NewCompetitionService(competitionRepository)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return server.HandlerFromMux(server.NewServer(competitionService), router)
	})
}
