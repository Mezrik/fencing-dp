package query

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type Group struct {
	ID            uuid.UUID `json:"id" ts_type:"UUID"`
	Name          string    `json:"name"`
	PisteNumber   *int64    `json:"pisteNumber"`
	CompetitionID uuid.UUID `json:"competitionId" ts_type:"UUID"`
}

func ToGroupQueryFromEntity(group *entities.Group) *Group {
	return &Group{
		ID:            group.ID,
		Name:          group.Name(),
		PisteNumber:   group.PisteNumber(),
		CompetitionID: group.CompetitionID(),
	}
}

func ToGroupQueryListFromEntities(groups []*entities.Group) []*Group {
	groupList := make([]*Group, 0, len(groups))
	for _, group := range groups {
		groupList = append(groupList, ToGroupQueryFromEntity(group))
	}
	return groupList
}
