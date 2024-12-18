package query

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type CompetitionParameters struct {
	ExpectedParticipants       int                     `json:"expectedParticipants"`
	DeploymentType             entities.DeploymentType `json:"deploymentType"`
	GroupHits                  int                     `json:"groupHits"`
	EliminationHits            int                     `json:"eliminationHits"`
	QualificationBasedOnRounds int                     `json:"qualificationBasedOnRounds"`
}

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
	Parameters      *CompetitionParameters       `json:"parameters"`
}

func ToCompetitionParametersFromEntity(parameters *entities.CompetitionParameters) *CompetitionParameters {
	if parameters == nil {
		return nil
	}

	return &CompetitionParameters{
		ExpectedParticipants:       parameters.ExpectedParticipants(),
		DeploymentType:             entities.DeploymentType(parameters.DeploymentType()),
		GroupHits:                  parameters.GroupHits(),
		EliminationHits:            parameters.EliminationHits(),
		QualificationBasedOnRounds: parameters.QualificationBasedOnRounds(),
	}
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

		Parameters: ToCompetitionParametersFromEntity(competition.Parameters()),
	}
}
func ToCompetitionQueryListFromEntities(competitions []*entities.Competition) []*Competition {
	competitionList := make([]*Competition, 0, len(competitions))
	for _, competition := range competitions {
		competitionList = append(competitionList, ToCompetitionQueryFromEntity(competition))
	}
	return competitionList
}
