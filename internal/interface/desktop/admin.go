package desktop

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/app/interfaces"
	"github.com/Mezrik/fencing-dp/internal/app/services"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/database/inmemory"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/database/inmemory/repositories"
)

type Admin struct {
	ctx                context.Context
	competitionService interfaces.CompetitionService
}

func New() *Admin {
	return &Admin{}
}

func (a *Admin) Startup(ctx context.Context) {
	a.ctx = ctx

	db, _ := inmemory.NewConnection()

	competitionRepository := repositories.NewInMemoryCompetitionRepository(db)

	a.competitionService = services.NewCompetitionService(competitionRepository)
}
