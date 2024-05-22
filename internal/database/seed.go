package database

import (
	"gorm.io/gorm"
	"mp.fencing.core/internal/models"
)

func Seed(db *gorm.DB) {
	// Create categories
	category1 := models.CompetitionCategory{Name: "Category 1"}
	category2 := models.CompetitionCategory{Name: "Category 2"}
	db.Create(&category1)
	db.Create(&category2)

	// Create weapons
	weapon1 := models.Weapon{Name: "Weapon 1"}
	weapon2 := models.Weapon{Name: "Weapon 2"}
	db.Create(&weapon1)
	db.Create(&weapon2)

	// Create competitions
	competition1 := models.Competition{
		Name:       "Competition 1",
		CategoryID: category1.ID,
		WeaponID:   weapon1.ID,
	}
	competition2 := models.Competition{
		Name:       "Competition 2",
		CategoryID: category2.ID,
		WeaponID:   weapon2.ID,
	}
	competition3 := models.Competition{
		Name:       "Competition 3",
		CategoryID: category1.ID,
		WeaponID:   weapon2.ID,
	}

	// Seed competitions
	db.Create(&competition1)
	db.Create(&competition2)
	db.Create(&competition3)
}
