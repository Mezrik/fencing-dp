package query

import (
	"github.com/Mezrik/fencing-dp/internal/match/domain/entities"
	"github.com/google/uuid"
)

type Match struct {
	ID               uuid.UUID `json:"id" ts_type:"UUID"`
	GroupID          uuid.UUID `json:"groupId" ts_type:"UUID"`
	ParticipantOneID uuid.UUID `json:"participantOneId" ts_type:"UUID"`
	ParticipantTwoID uuid.UUID `json:"participantTwoId" ts_type:"UUID"`
}

type MatchDetail struct {
	ID               uuid.UUID     `json:"id" ts_type:"UUID"`
	GroupID          uuid.UUID     `json:"groupId" ts_type:"UUID"`
	ParticipantOneID uuid.UUID     `json:"participantOneId" ts_type:"UUID"`
	ParticipantTwoID uuid.UUID     `json:"participantTwoId" ts_type:"UUID"`
	State            []*MatchState `json:"state"`
}

func ToMatchQueryFromEntity(m *entities.Match) *Match {
	return &Match{
		ID:               m.ID,
		GroupID:          m.GroupID(),
		ParticipantOneID: m.ParticipantOneID(),
		ParticipantTwoID: m.ParticipantTwoID(),
	}
}

func ToMatchQueryListFromEntities(m []*entities.Match) []*Match {
	matchList := make([]*Match, 0, len(m))
	for _, v := range m {
		matchList = append(matchList, ToMatchQueryFromEntity(v))
	}
	return matchList
}

func ToMatchDetailQueryFromEntity(m *entities.Match) *MatchDetail {
	return &MatchDetail{
		ID:               m.ID,
		GroupID:          m.GroupID(),
		ParticipantOneID: m.ParticipantOneID(),
		ParticipantTwoID: m.ParticipantTwoID(),
		State:            ToMatchStateQueryListFromEntities(m.State()),
	}
}

func ToMatchDetailQueryListFromEntities(m []*entities.Match) []*MatchDetail {
	matchList := make([]*MatchDetail, 0, len(m))
	for _, v := range m {
		matchList = append(matchList, ToMatchDetailQueryFromEntity(v))
	}
	return matchList
}
