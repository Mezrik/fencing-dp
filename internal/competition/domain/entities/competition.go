package entities

import (
	"errors"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type CompetitionTypeEnum string

const (
	National      CompetitionTypeEnum = "national"
	International CompetitionTypeEnum = "international"
)

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
	Mixed  GenderEnum = "mixed"
)

type Competition struct {
	common.Entity
	name            string
	organizerName   string
	federationName  string
	competitionType CompetitionTypeEnum
	category        CompetitionCategory
	gender          GenderEnum
	weapon          Weapon
	date            time.Time
}

func (g GenderEnum) TSName() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	case Mixed:
		return "mixed"
	default:
		return "???"
	}
}

func (ct CompetitionTypeEnum) TSName() string {
	switch ct {
	case National:
		return "national"
	case International:
		return "international"
	default:
		return "???"
	}
}

func NewCompetition(name string, orgName string, fedName string, comptType CompetitionTypeEnum, category CompetitionCategory, gender GenderEnum, weapon Weapon, date time.Time) (*Competition, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &Competition{
		Entity:          common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		name:            name,
		organizerName:   orgName,
		federationName:  fedName,
		competitionType: comptType,
		category:        category,
		gender:          gender,
		weapon:          weapon,
		date:            date,
	}, nil
}

func UnmarshalCompetition(id uuid.UUID, createdAt time.Time, updatedAt *time.Time, name string, orgName string, fedName string, comptType CompetitionTypeEnum, category CompetitionCategory, gender GenderEnum, weapon Weapon, date time.Time) *Competition {
	return &Competition{
		Entity:          common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		name:            name,
		organizerName:   orgName,
		federationName:  fedName,
		competitionType: comptType,
		category:        category,
		gender:          gender,
		weapon:          weapon,
		date:            date,
	}
}

func (c Competition) Name() string {
	return c.name
}

func (c Competition) OrganizerName() string {
	return c.organizerName
}

func (c Competition) FederationName() string {
	return c.federationName
}

func (c Competition) CompetitionType() CompetitionTypeEnum {
	return c.competitionType
}

func (c Competition) Category() CompetitionCategory {
	return c.category
}

func (c Competition) Gender() GenderEnum {
	return c.gender
}

func (c Competition) Weapon() Weapon {
	return c.weapon
}

func (c Competition) Date() time.Time {
	return c.date
}
