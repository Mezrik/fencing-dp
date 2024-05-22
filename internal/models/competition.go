package models

import "gorm.io/gorm"

type Competition struct {
	gorm.Model
	Name       string `json:"name"`
	Category   CompetitionCategory
	CategoryID uint
	Weapon     Weapon
	WeaponID   uint
}
