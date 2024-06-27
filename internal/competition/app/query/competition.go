package query

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type Competition struct {
	ID   uuid.UUID
	Name string `json:"name"`
}

func ToCompetitionQueryFromEntity(competition *entities.Competition) *Competition {
	return &Competition{
		ID:   competition.ID,
		Name: competition.Name(),
	}
}
func ToCompetitionQueryListFromEntities(competitions []*entities.Competition) []*Competition {
	competitionList := make([]*Competition, 0, len(competitions))
	for _, competition := range competitions {
		competitionList = append(competitionList, ToCompetitionQueryFromEntity(competition))
	}
	return competitionList
}
