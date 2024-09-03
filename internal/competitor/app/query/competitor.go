package query

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
)

type Competitor struct {
	Surname    string    `json:"surname"`
	Firstname  string    `json:"firstname"`
	Club       Club      `json:"club"`
	Gender     string    `json:"gender"`
	License    string    `json:"license"`
	LicenseFie *string   `json:"licenseFie"`
	Birthdate  time.Time `json:"birthdate"`
}

func ToCompetitorQueryFromEntity(c *entities.Competitor) *Competitor {
	return &Competitor{
		Surname:    c.Surname(),
		Firstname:  c.FirstName(),
		Club:       Club{ID: c.Club().ID, Name: c.Club().Name()},
		Gender:     string(c.Gender()),
		License:    c.License(),
		LicenseFie: c.LicenseFie(),
		Birthdate:  c.Birthdate(),
	}
}

func ToCompetitorQueryListFromEntities(c []*entities.Competitor) []*Competitor {
	competitors := make([]*Competitor, 0, len(c))

	for _, v := range c {
		competitors = append(competitors, ToCompetitorQueryFromEntity(v))
	}
	return competitors
}
