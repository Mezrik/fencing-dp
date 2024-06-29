package inmemory

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

func NewConnection(logger *logrus.Entry) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "/desktop-sqlite.db")

	if err != nil {
		logger.WithError(err).Error("failed to connect database")
	}

	// TODO: do the auto migration

	return db, err
}
