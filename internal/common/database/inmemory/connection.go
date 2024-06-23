package inmemory

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func NewConnection() (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// models := []interface{}{
	// 	&models.CompetitionModel{},
	// 	&models.WeaponModel{},
	// 	&models.CompetitionCategoryModel{},
	// }

	// for _, m := range models {
	// 	if err := gormDB.AutoMigrate(m); err != nil {
	// 		// TODO: handle error
	// 	}
	// }

	gormDB.AutoMigrate()

	return gormDB, err
}
