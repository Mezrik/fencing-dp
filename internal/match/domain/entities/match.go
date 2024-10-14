package entities

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type Match struct {
	common.Entity
	groupID          uuid.UUID
	participantOneID uuid.UUID
	participantTwoID uuid.UUID
	state            []*MatchState
}

func NewMatch(groupId uuid.UUID, participantOneId uuid.UUID, participantTwoId uuid.UUID) *Match {
	return &Match{
		Entity:           common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		groupID:          groupId,
		participantOneID: participantOneId,
		participantTwoID: participantTwoId,

		state: make([]*MatchState, 0),
	}
}

func UnmarshalMatch(id uuid.UUID, groupId uuid.UUID, participantOneId uuid.UUID, participantTwoId uuid.UUID, state []*MatchState, createdAt time.Time, updatedAt *time.Time) *Match {
	return &Match{
		Entity:           common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		groupID:          groupId,
		participantOneID: participantOneId,
		participantTwoID: participantTwoId,
		state:            state,
	}
}

func (m *Match) GroupID() uuid.UUID {
	return m.groupID
}

func (m *Match) ParticipantOneID() uuid.UUID {
	return m.participantOneID
}

func (m *Match) ParticipantTwoID() uuid.UUID {
	return m.participantTwoID
}

func (m *Match) State() []*MatchState {
	return m.state
}
