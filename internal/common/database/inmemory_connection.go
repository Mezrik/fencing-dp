package database

import (
	"context"

	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

func mapper(s string) string { return s }

func NewConnection(logger *logrus.Entry, ctx context.Context, dbMigrations embed.FS) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "sqlite3", "file:test.db")

	if err != nil {
		logger.WithError(err).Error("failed to connect database")
	}

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "desktop-sqlite3", // TODO: This should be configurable
	}

	n, err := migrate.Exec(db.DB, "sqlite3", migrations, migrate.Up)

	if err != nil {
		logger.WithError(err).Error("failed to migrate database")
	}

	logger.Infof("Applied %d migrations", n)

	db.MapperFunc(mapper)

	return db, err
}
