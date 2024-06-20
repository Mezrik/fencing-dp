package common

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
