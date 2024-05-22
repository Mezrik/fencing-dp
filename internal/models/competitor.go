package models

import "gorm.io/gorm"

type Competitor struct {
	gorm.Model
	CompetionID uint
	FighterID   uint
	Competion   Competition
	Fighter     Fighter
}
