package models

import "github.com/google/uuid"

type CompetitionModel struct {
	Model
	Name       string `json:"name"`
	Category   CompetitionCategoryModel
	CategoryID uuid.UUID
	Weapon     WeaponModel
	WeaponID   uuid.UUID
}
