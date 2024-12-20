package entities

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type Group struct {
	common.Entity
	name           string
	piste_number   *int64
	competition_id uuid.UUID
}

func NewCompetitionGroup(name string, competitionId uuid.UUID) *Group {
	return &Group{
		Entity:         common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		name:           name,
		competition_id: competitionId,
	}
}

func UnmarshalCompetitionGroup(id uuid.UUID, name string, piste_number *int64, createdAt time.Time, updatedAt *time.Time, competitionId uuid.UUID) *Group {
	return &Group{
		Entity:         common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		name:           name,
		competition_id: competitionId,
		piste_number:   piste_number,
	}
}

func (c Group) Name() string {
	return c.name
}

func (c Group) PisteNumber() *int64 {
	return c.piste_number
}

func (c Group) CompetitionID() uuid.UUID {
	return c.competition_id
}
