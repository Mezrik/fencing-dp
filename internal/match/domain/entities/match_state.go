package entities

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type MatchState struct {
	common.Entity
	matchId uuid.UUID
	change  MatchChangeEnum
	pointTo uuid.NullUUID
}

type MatchChangeEnum string

const (
	MatchStart      = "match_start"
	FightStart      = "fight_start"
	FightStop       = "fight_stop"
	PointAdded      = "point_added"
	PointSubtracted = "point_subtracted"
	MatchEnd        = "match_end"
)

func (g MatchChangeEnum) TSName() string {
	switch g {
	case MatchStart:
		return "match_start"
	case FightStart:
		return "fight_start"
	case FightStop:
		return "fight_stop"
	case PointAdded:
		return "point_added"
	case PointSubtracted:
		return "point_subtracted"
	case MatchEnd:
		return "match_end"
	default:
		return "???"
	}
}

func NewMatchState(matchId uuid.UUID, change MatchChangeEnum, pointTo uuid.NullUUID) *MatchState {
	return &MatchState{
		Entity:  common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		matchId: matchId,
		change:  change,
		pointTo: pointTo,
	}
}

func UnmarshalMatchState(id uuid.UUID, matchId uuid.UUID, change MatchChangeEnum, pointTo uuid.NullUUID, createdAt time.Time, updatedAt *time.Time) *MatchState {
	return &MatchState{
		Entity:  common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		matchId: matchId,
		change:  change,
		pointTo: pointTo,
	}
}

func (m *MatchState) MatchId() uuid.UUID {
	return m.matchId
}

func (m *MatchState) Change() MatchChangeEnum {
	return m.change
}

func (m *MatchState) PointTo() uuid.NullUUID {
	return m.pointTo
}
