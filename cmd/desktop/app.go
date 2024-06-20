package main

import (
	"context"
	"fmt"

	"github.com/Mezrik/fencing-dp/internal/app/common"
	"github.com/Mezrik/fencing-dp/internal/app/interfaces"
	"github.com/Mezrik/fencing-dp/internal/app/services"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/database/inmemory"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/database/inmemory/repositories"
)

// App struct
type App struct {
	ctx                context.Context
	competitionService interfaces.CompetitionService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	db, _ := inmemory.NewConnection()

	competitionRepository := repositories.NewInMemoryCompetitionRepository(db)

	a.competitionService = services.NewCompetitionService(competitionRepository)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetCompetitions() []*common.CompetitionResult {

	competitions, _ := a.competitionService.GetAllCompetitions()

	return competitions
}
