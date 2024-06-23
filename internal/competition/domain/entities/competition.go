package entities

import (
	"errors"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type CompetitionType int

const (
	National CompetitionType = iota + 1
	International
)

type Gender int

const (
	Male Gender = iota + 1
	Female
)

type Competition struct {
	common.Entity
	Name           string
	OrganizerName  string
	FederationName string
	Type           CompetitionType
	Category       CompetitionCategory
	Gender         Gender
	Weapon         Weapon
	Date           time.Time
}

func (c *Competition) validate() error {
	if c.Name == "" {
		return errors.New("name must not be empty")
	}

	return nil
}

func NewCompetition(name string) (*Competition, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &Competition{
		Entity: common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:   name,
	}, nil
}

func (c *Competition) UpdateName(name string) error {
	c.Name = name
	c.UpdatedAt = time.Now()

	return c.validate()
}

func (c *Competition) UpdateCategory(category CompetitionCategory) error {
	c.Category = category
	c.UpdatedAt = time.Now()

	return c.validate()
}

type ValidatedCompetition struct {
	Competition
	Errors []string
}

func (vc *ValidatedCompetition) IsValid() bool {
	return len(vc.Errors) == 0
}
