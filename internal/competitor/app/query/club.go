package query

import (
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/google/uuid"
)

type Club struct {
	ID   uuid.UUID `json:"id" ts_type:"UUID"`
	Name string    `json:"name"`
}

func ToClubQueryFromEntity(weapon *entities.Club) *Club {
	return &Club{
		ID:   weapon.ID,
		Name: weapon.Name(),
	}
}
