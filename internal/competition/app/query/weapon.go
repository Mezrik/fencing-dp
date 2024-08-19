package query

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type Weapon struct {
	ID   uuid.UUID `json:"id" ts_type:"UUID"`
	Name string    `json:"name"`
}

func ToWeaponQueryFromEntity(weapon *entities.Weapon) *Weapon {
	return &Weapon{
		ID:   weapon.ID,
		Name: weapon.Name(),
	}
}
func ToWeaponQueryListFromEntities(weapons []*entities.Weapon) []*Weapon {
	weaponsList := make([]*Weapon, 0, len(weapons))
	for _, weapon := range weapons {
		weaponsList = append(weaponsList, ToWeaponQueryFromEntity(weapon))
	}
	return weaponsList
}
