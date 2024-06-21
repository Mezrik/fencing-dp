package main

import (
	"context"
	"fmt"

	"github.com/Mezrik/fencing-dp/internal/interface/desktop"
)

// App struct
type App struct {
	ctx   context.Context
	admin *desktop.Admin
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context, admin *desktop.Admin) {
	a.ctx = ctx
	a.admin = admin
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
