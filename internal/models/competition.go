package models

import "gorm.io/gorm"

type Competition struct {
	gorm.Model
	Name       string
	Category   CompetitionCategory
	CategoryID uint
	Weapon     Weapon
	WeaponID   uint
}
