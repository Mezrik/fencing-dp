package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mp.fencing.core/internal/core"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	core.NewApp(db)
}
