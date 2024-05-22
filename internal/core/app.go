package core

import (
	"gorm.io/gorm"
	"mp.fencing.core/internal/database"
	"mp.fencing.core/internal/models"
)

type App struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) *App {
	models := []interface{}{
		&models.Competition{},
		&models.Competitor{},
		&models.Fighter{},
		&models.Weapon{},
		&models.CompetitionCategory{},
	}

	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			// TODO: handle error
		}
	}

	database.Seed(db)

	return &App{
		db: db,
	}
}
