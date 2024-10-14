package main

import (
	"context"

	"github.com/Mezrik/fencing-dp/frontend"
	competitionEntities "github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	competitorEntities "github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/interface/desktop"
	matchEntities "github.com/Mezrik/fencing-dp/internal/match/domain/entities"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	// Create an instance of the app structure
	app := NewApp()
	admin := desktop.New()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Fencing",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: frontend.Assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			admin.Startup(ctx)
			app.Startup(ctx, admin)
		},
		Bind: []interface{}{
			app,
			admin,
		},
		EnumBind: []interface{}{
			[]competitionEntities.GenderEnum{"male", "female", "mixed"},
			[]competitionEntities.CompetitionTypeEnum{"national", "international"},
			[]competitorEntities.GenderEnum{"male", "female", "mixed"},
			[]matchEntities.MatchChangeEnum{"match_start", "match_end", "fight_start", "fight_stop", "point_added", "point_subtracted"},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
