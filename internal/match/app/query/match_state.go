package query

import (
	"github.com/Mezrik/fencing-dp/internal/match/domain/entities"
	"github.com/google/uuid"
)

type MatchState struct {
	ID      uuid.UUID                `json:"id" ts_type:"UUID"`
	Change  entities.MatchChangeEnum `json:"change"`
	PointTo uuid.NullUUID            `json:"pointTo"`
}

func ToMatchStateQueryFromEntity(m *entities.MatchState) *MatchState {
	return &MatchState{
		ID:      m.ID,
		Change:  m.Change(),
		PointTo: m.PointTo(),
	}
}

func ToMatchStateQueryListFromEntities(m []*entities.MatchState) []*MatchState {
	matchStateList := make([]*MatchState, 0, len(m))
	for _, v := range m {
		matchStateList = append(matchStateList, ToMatchStateQueryFromEntity(v))
	}
	return matchStateList
}
