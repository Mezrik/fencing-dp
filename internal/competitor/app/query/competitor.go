package query

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/google/uuid"
)

type Competitor struct {
	ID             uuid.UUID           `json:"id" ts_type:"UUID"`
	Surname        string              `json:"surname"`
	Firstname      string              `json:"firstname"`
	Club           *Club               `json:"club"`
	Gender         entities.GenderEnum `json:"gender"`
	License        *string             `json:"license"`
	LicenseFie     *string             `json:"licenseFie"`
	Birthdate      *time.Time          `json:"birthdate"`
	HasMissingInfo bool                `json:"hasMissingInfo"`
}

func ToCompetitorQueryFromEntity(c *entities.Competitor) *Competitor {
	var club *Club

	if c.Club() != nil {
		club = &Club{ID: c.Club().ID, Name: c.Club().Name()}
	}

	var birthdate *time.Time

	if c.Birthdate() != nil && !c.Birthdate().IsZero() {
		birthdate = c.Birthdate()
	}

	return &Competitor{
		ID:             c.ID,
		Surname:        c.Surname(),
		Firstname:      c.FirstName(),
		Club:           club,
		Gender:         c.Gender(),
		License:        c.License(),
		LicenseFie:     c.LicenseFie(),
		Birthdate:      birthdate,
		HasMissingInfo: c.HasMissingInfo(),
	}
}

func ToCompetitorQueryListFromEntities(c []*entities.Competitor) []*Competitor {
	competitors := make([]*Competitor, 0, len(c))

	for _, v := range c {
		competitors = append(competitors, ToCompetitorQueryFromEntity(v))
	}
	return competitors
}
