package migrations

import (
	"embed"
)

//go:embed desktop-sqlite3/*.sql
var SQLiteMigrations embed.FS
