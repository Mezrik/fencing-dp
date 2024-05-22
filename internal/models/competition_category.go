package models

import "gorm.io/gorm"

type CompetitionCategory struct {
	gorm.Model
	Name string
}
