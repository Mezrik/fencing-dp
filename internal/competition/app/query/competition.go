package query

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type Competition struct {
	ID              uuid.UUID                    `json:"id" ts_type:"UUID"`
	Name            string                       `json:"name"`
	OrganizerName   string                       `json:"organizerName"`
	FederationName  string                       `json:"federationName"`
	CompetitionType entities.CompetitionTypeEnum `json:"competitionType"`
	Category        CompetitionCategory          `json:"category"`
	Gender          entities.GenderEnum          `json:"gender"`
	Weapon          Weapon                       `json:"weapon"`
	Date            time.Time                    `json:"date"`
}

func ToCompetitionQueryFromEntity(competition *entities.Competition) *Competition {
	return &Competition{
		ID:   competition.ID,
		Name: competition.Name(),

		OrganizerName:   competition.OrganizerName(),
		FederationName:  competition.FederationName(),
		CompetitionType: competition.CompetitionType(),

		Category: CompetitionCategory{
			ID:   competition.Category().ID,
			Name: competition.Category().Name(),
		},

		Gender: entities.GenderEnum(competition.Gender()),

		Weapon: Weapon{
			ID:   competition.Weapon().ID,
			Name: competition.Weapon().Name(),
		},

		Date: competition.Date(),
	}
}
func ToCompetitionQueryListFromEntities(competitions []*entities.Competition) []*Competition {
	competitionList := make([]*Competition, 0, len(competitions))
	for _, competition := range competitions {
		competitionList = append(competitionList, ToCompetitionQueryFromEntity(competition))
	}
	return competitionList
}
