package database

import (
	"github.com/Mezrik/fencing-dp/internal/infrastructure/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Create categories
	category1 := models.CompetitionCategoryModel{Model: models.Model{ID: uuid.New()}, Name: "Category 1"}
	category2 := models.CompetitionCategoryModel{Model: models.Model{ID: uuid.New()}, Name: "Category 2"}
	db.Create(&category1)
	db.Create(&category2)

	// Create weapons
	weapon1 := models.WeaponModel{
		Model: models.Model{ID: uuid.New()},
		Name:  "Weapon 1"}
	weapon2 := models.WeaponModel{Model: models.Model{ID: uuid.New()}, Name: "Weapon 2"}
	db.Create(&weapon1)
	db.Create(&weapon2)

	// Create competitions
	competition1 := models.CompetitionModel{
		Model:      models.Model{ID: uuid.New()},
		Name:       "Competition 1",
		CategoryID: category1.ID,
		WeaponID:   weapon1.ID,
	}
	competition2 := models.CompetitionModel{
		Model:      models.Model{ID: uuid.New()},
		Name:       "Competition 2",
		CategoryID: category2.ID,
		WeaponID:   weapon2.ID,
	}
	competition3 := models.CompetitionModel{
		Model:      models.Model{ID: uuid.New()},
		Name:       "Competition 3",
		CategoryID: category1.ID,
		WeaponID:   weapon2.ID,
	}

	// Seed competitions
	db.Create(&competition1)
	db.Create(&competition2)
	db.Create(&competition3)
}
