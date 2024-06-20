package inmemory

import (
	"log"

	"github.com/Mezrik/fencing-dp/internal/infrastructure/database"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/models"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func NewConnection() (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	models := []interface{}{
		&models.CompetitionModel{},
		&models.WeaponModel{},
		&models.CompetitionCategoryModel{},
	}

	for _, m := range models {
		if err := gormDB.AutoMigrate(m); err != nil {
			// TODO: handle error
		}
	}

	gormDB.AutoMigrate()

	database.Seed(gormDB)

	return gormDB, err
}
