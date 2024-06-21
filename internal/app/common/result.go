package common

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	Id        uuid.UUID `json:"id" ts_type:"UUID"`
	CreatedAt time.Time `json:"createdAt" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	UpdatedAt time.Time `json:"updatedAt" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
}
