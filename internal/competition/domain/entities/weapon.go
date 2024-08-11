package entities

import (
	"errors"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type Weapon struct {
	common.Entity
	name string
}

func NewWeapon(name string) (*Weapon, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &Weapon{
		Entity: common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		name:   name,
	}, nil
}

func UnmarshalWeapon(id uuid.UUID, name string, createdAt time.Time, updatedAt *time.Time) *Weapon {
	return &Weapon{
		Entity: common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: nil},
		name:   name,
	}
}

func (w Weapon) Name() string {
	return w.name
}
