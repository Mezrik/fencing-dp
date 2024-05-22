package models

import "gorm.io/gorm"

type Fighter struct {
	gorm.Model
	Name string
}
