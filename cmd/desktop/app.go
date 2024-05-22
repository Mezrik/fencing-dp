package main

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mp.fencing.core/internal/core"
	"mp.fencing.core/internal/models"
)

// App struct
type App struct {
	ctx  context.Context
	core *core.Core
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	a.core = core.SetupCore(db)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetCompetitions() []models.Competition {
	var competitions []models.Competition
	a.core.DB.Find(&competitions)
	return competitions
}
