package models

import (
	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type CompetitionModel struct {
	common.Model
	Name       string
	Category   CompetitionCategoryModel
	CategoryID uuid.UUID
	Weapon     WeaponModel
	WeaponID   uuid.UUID
}
