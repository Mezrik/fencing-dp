package models

import "gorm.io/gorm"

type Weapon struct {
	gorm.Model
	Name string
}
