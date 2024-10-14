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

func GetInt64Ptr(nt sql.NullInt64) *int64 {
	if nt.Valid {
		return &nt.Int64
	}
	return nil
}
