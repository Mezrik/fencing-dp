package core

import (
	"gorm.io/gorm"
	"mp.fencing.core/internal/database"
	"mp.fencing.core/internal/models"
)

type Core struct {
	DB *gorm.DB
}

func SetupCore(db *gorm.DB) *Core {
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

	return &Core{
		DB: db,
	}
}
