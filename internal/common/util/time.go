package util

import (
	"database/sql"
	"time"
)

func GetTimePtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
