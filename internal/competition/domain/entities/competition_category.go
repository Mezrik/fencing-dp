package entities

import (
	"errors"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type CompetitionCategory struct {
	common.Entity
	name string
}

func NewCompetitionCategory(name string) (*CompetitionCategory, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &CompetitionCategory{
		Entity: common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		name:   name,
	}, nil
}

func UnmarshalCompetitionCategory(id uuid.UUID, name string, createdAt time.Time, updatedAt *time.Time) *CompetitionCategory {
	return &CompetitionCategory{
		Entity: common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		name:   name,
	}
}

func (c CompetitionCategory) Name() string {
	return c.name
}
