package entities

import (
	"errors"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type Club struct {
	common.Entity
	name string
}

func NewClub(name string) (*Club, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}
	return &Club{
		Entity: common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		name:   name,
	}, nil
}

func UnmarshalClub(id uuid.UUID, name string, createdAt time.Time, updatedAt *time.Time) *Club {
	return &Club{
		Entity: common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		name:   name,
	}
}

func (c *Club) Name() string {
	return c.name
}
